package torrent

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
)

type ProfileUpdateRequest struct {
	TorrentID int32 `json:"torrent_id" binding:"required"`
}

// Delete 种子删除 (POST /delete)
func Delete(c *gin.Context) {
	// 参数绑定
	var req ProfileUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	err := db.DeleteTorrent(req.TorrentID)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordDeleteFailed, err.Error())
		log.Logger.Errorf("delete torrent failed, torrent_id: %d, error: %s", req.TorrentID, err.Error())
		return
	}

	resp.OK(c)
	cache.ClearTorrentDetailCacheByTorrentID(req.TorrentID)

	log.Logger.Infof("delete torrent success, torrent_id: %d", req.TorrentID)
}
