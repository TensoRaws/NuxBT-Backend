package torrent

import (
	"bytes"
	"mime/multipart"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/oss"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/role"
	"github.com/TensoRaws/NuxBT-Backend/module/torrent"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

type UploadRequest struct {
	File        multipart.FileHeader `form:"torrent_file" type:"blob" binding:"required"`
	AnidbID     int32                `form:"anidb_id" binding:"required"`
	AudioCodec  string               `form:"audio_codec" binding:"required,oneof=FLAC AAC AC3 DTS DDP LPCM other"`
	Description string               `form:"description" binding:"required"`
	Essay       *string              `form:"essay" binding:"omitempty"`
	Genre       string               `form:"genre" binding:"required,oneof=BDrip WEBrip DVDrip Remux Blu-ray WEB-DL DVD HDTV other"` //nolint:lll
	Img         string               `form:"img" binding:"required"`
	Language    string               `form:"language" binding:"required,oneof=Chinese English Japanese other"`
	Official    string               `form:"official" binding:"required,oneof=true false"`
	Resolution  string               `form:"resolution" binding:"required,oneof=480p 720p 1080p 2160p other"`
	Subtitle    string               `form:"subtitle" binding:"required"`
	Title       string               `form:"title" binding:"required"`
	VideoCodec  string               `form:"video_codec" binding:"required,oneof=H.265 H.264 AV1 VP9 other"`
}

// Upload 上传种子 (POST /upload)
func Upload(c *gin.Context) {
	// 绑定参数
	var req UploadRequest
	if err := c.ShouldBind(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	userID, _ := resp.GetUserIDFromGinContext(c)

	isOfficial := req.Official == "true"

	roles, err := resp.GetRolesFromGinContext(c)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		log.Logger.Error("failed to get roles from gin context: " + err.Error())
		return
	}

	// 判断是否有官方权限，没有的话不能发官种
	if !util.CheckStringInSlice(role.ADMIN, roles) && isOfficial {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, "official permission required")
		log.Logger.Errorf("official permission required, user ID: %v", userID)
		return
	}

	// 判断发布权限，是否可以直接发种
	var status string
	if util.CheckStringInSlice(role.UPLOADER, roles) || util.CheckStringInSlice(role.ADMIN, roles) {
		status = STATUS_APPROVED
	} else {
		status = STATUS_PENDING
	}

	// 开始解析种子文件
	torrentFile, err := torrent.NewBitTorrentFileFromMultipart(&req.File)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		log.Logger.Error("failed to open file: " + err.Error())
		return
	}

	// 开始清洗种子，官种和非官种的清洗策略不同，官种可能会变更 hash
	var strategy torrent.BitTorrentFileEditStrategy
	if isOfficial {
		strategy = torrent.BitTorrentFileEditStrategy{
			AnnounceList: torrent.TRACKER_LIST,
			Comment:      &config.ServerConfig.Name,
			InfoSource:   &config.ServerConfig.Name,
		}
	} else {
		strategy = torrent.BitTorrentFileEditStrategy{
			AnnounceList: torrent.TRACKER_LIST,
			Comment:      &config.ServerConfig.Name,
		}
	}
	err = torrentFile.Repack(&strategy)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		log.Logger.Error("failed to repack torrent file: " + err.Error())
		return
	}

	// 计算种子文件的哈希值，大小，结构
	hash := torrentFile.GetHash()
	size := torrentFile.GetTotalSize()
	filelist := torrentFile.GetFileList()

	// 上传到 OSS
	torrentBytes, err := torrentFile.ConvertToBytes()
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		log.Logger.Error("failed to convert torrent file to bytes: " + err.Error())
		return
	}
	torrentBytesReader := bytes.NewReader(torrentBytes)
	torrentKey := torrentFile.Info.Name + "--" + hash + ".torrent"
	err = oss.Put(torrentKey, torrentBytesReader)
	if err != nil {
		resp.AbortWithMsg(c, code.OssErrorPutFailed, err.Error())
		log.Logger.Error("failed to put torrent file to oss: " + err.Error())
		return
	}

	// 保存到数据库
	err = db.CreateTorrent(&model.Torrent{
		Hash:        hash,
		UploaderID:  userID,
		Official:    isOfficial,
		Size:        size,
		Status:      status,
		Title:       req.Title,
		Subtitle:    req.Subtitle,
		Essay:       *req.Essay,
		Description: req.Description,
		Genre:       req.Genre,
		AnidbID:     req.AnidbID,
		Img:         req.Img,
		Resolution:  req.Resolution,
		VideoCodec:  req.VideoCodec,
		AudioCodec:  req.AudioCodec,
		Language:    req.Language,
		URL:         torrentKey,
		FileList:    util.StructToString(filelist),
	})
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordCreateFailed, err.Error())
		log.Logger.Error("failed to create torrent record: " + err.Error())
		return
	}

	// 从数据库获取上传的种子
	newTorrent, err := db.GetTorrentByHash(hash)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, err.Error())
		log.Logger.Error("failed to get torrent record: " + err.Error())
		return
	}

	torrentInfo, err := GetTorrentInfo(newTorrent)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		log.Logger.Error("failed to get torrent info: " + err.Error())
		return
	}

	resp.OKWithData(c, torrentInfo)
	log.Logger.Infof("upload torrent success, user ID: %v", userID)
}
