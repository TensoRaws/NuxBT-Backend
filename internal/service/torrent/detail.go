package torrent

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
)

type DetailRequest struct {
	TorrentID int32 `form:"torrent_id" binding:"required"`
}

// Detail 获取种子文件列表 (GET /detail)
func Detail(c *gin.Context) {
	// 绑定参数
	var req DetailRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	bt, err := db.GetTorrentByID(req.TorrentID)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, err.Error())
		return
	}

	torrentInfo, err := GetTorrentInfo(bt)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		log.Logger.Error("failed to get torrent info: " + err.Error())
		return
	}

	resp.OKWithData(c, torrentInfo)
	log.Logger.Infof("get torrent detail success: %v", req.TorrentID)
}
