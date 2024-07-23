package torrent

import (
	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/role"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

type EditRequest struct {
	TorrentID   int32  `form:"torrent_id" binding:"required"`
	AnidbID     int32  `form:"anidb_id" binding:"required"`
	AudioCodec  string `form:"audio_codec" binding:"required,oneof=FLAC AAC AC3 DTS DDP LPCM other"`
	Description string `form:"description" binding:"required"`
	Essay       string `form:"essay" binding:"required"`
	Genre       string `form:"genre" binding:"required,oneof=BDrip WEBrip DVDrip Remux Blu-ray WEB-DL DVD HDTV other"` //nolint:lll
	Img         string `form:"img" binding:"required"`
	Language    string `form:"language" binding:"required,oneof=Chinese English Japanese other"`
	Resolution  string `form:"resolution" binding:"required,oneof=480p 720p 1080p 2160p other"`
	Subtitle    string `form:"subtitle" binding:"required"`
	Title       string `form:"title" binding:"required"`
	VideoCodec  string `form:"video_codec" binding:"required,oneof=H.265 H.264 AV1 VP9 other"`
}

// Edit 更新种子信息 (POST /edit)
func Edit(c *gin.Context) {
	// 参数绑定
	var req EditRequest
	if err := c.ShouldBind(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	userID, _ := resp.GetUserIDFromGinContext(c)

	roles, err := resp.GetRolesFromGinContext(c)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		log.Logger.Error("failed to get roles from gin context: " + err.Error())
		return
	}

	bt, err := db.GetTorrentByID(req.TorrentID)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, err.Error())
		log.Logger.Error("failed to get torrent by id: " + err.Error())
		return
	}

	if bt.UploaderID != userID && !util.CheckStringInSlice(role.ADMIN, roles) {
		resp.AbortWithMsg(c, code.AuthErrorNoPermission, "permission denied")
		log.Logger.Errorf("permission denied, user id: %v", userID)
		return
	}

	err = db.PatchTorrent(req.TorrentID, &model.Torrent{
		AnidbID:     req.AnidbID,
		AudioCodec:  req.AudioCodec,
		Description: req.Description,
		Essay:       req.Essay,
		Genre:       req.Genre,
		Img:         req.Img,
		Language:    req.Language,
		Resolution:  req.Resolution,
		Subtitle:    req.Subtitle,
		Title:       req.Title,
		VideoCodec:  req.VideoCodec,
	})
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordPatchFailed, err.Error())
		return
	}

	resp.OK(c)

	log.Logger.Infof("update torrent info success, torrent id: %v", req.TorrentID)
}
