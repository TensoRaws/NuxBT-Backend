package torrent

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/torrent"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
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

	var fileList FileListResponse
	err = util.StringToStruct(bt.FileList, &fileList)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		return
	}

	resp.OKWithData(c, &fileList)
	log.Logger.Infof("get fileList success, torrent ID: %v", req.TorrentID)
}
