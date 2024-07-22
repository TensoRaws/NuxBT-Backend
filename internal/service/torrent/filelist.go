package torrent

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/oss"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/torrent"
	"github.com/gin-gonic/gin"
)

type FileListRequest struct {
	TorrentID int32 `form:"torrent_id" binding:"required"`
}

type FileListResponse []torrent.BitTorrentFileListItem

// FileList 获取种子文件列表
func FileList(c *gin.Context) {
	// 绑定参数
	var req FileListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	bt, err := db.GetTorrentByID(req.TorrentID)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, err.Error())
		return
	}

	torrentBytes, err := oss.GetBytes(bt.URL)
	if err != nil {
		resp.AbortWithMsg(c, code.OssErrorGetFailed, err.Error())
		return
	}

	torrentFile, err := torrent.NewBitTorrentFileFromBytes(torrentBytes)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		return
	}

	fileList := torrentFile.GetFileList()

	resp.OKWithData(c, &fileList)
	log.Logger.Infof("get fileList success, torrent ID: %v", req.TorrentID)
}
