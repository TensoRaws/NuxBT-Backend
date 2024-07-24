package torrent

import (
	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
)

type ReviewRequest struct {
	TorrentID int32  `json:"torrent_id" binding:"required"`
	Status    string `json:"status" binding:"required,oneof= pending approved rejected"`
}

// Review 种子审核 (POST /review)
func Review(c *gin.Context) {
	// 参数绑定
	var req ReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	bt, err := db.GetTorrentByID(req.TorrentID)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, err.Error())
		log.Logger.Errorf("get torrent by id failed, torrent_id: %d, err: %v", req.TorrentID, err)
		return
	}

	if bt.Status == db.STATUS_APPROVED {
		resp.AbortWithMsg(c, code.AuthErrorNoPermission, "torrent already approved")
		log.Logger.Errorf("torrent already approved, torrent_id: %d", req.TorrentID)
		return
	}

	err = db.PatchTorrent(req.TorrentID, &model.Torrent{Status: req.Status})
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordPatchFailed, err.Error())
		return
	}

	resp.OK(c)

	log.Logger.Infof("review torrent success, torrent_id: %d, status: %s", req.TorrentID, req.Status)
}
