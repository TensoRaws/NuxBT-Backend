package torrent

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

type OfficialRequest struct {
	Order   string  `form:"order" binding:"required,oneof=asc desc"`
	OrderBy string  `form:"order_by" binding:"required,oneof=date size"`
	Page    int     `form:"page" binding:"required,min=1"`
	PerPage int     `form:"per_page" binding:"required,min=10"`
	Search  *string `form:"search" binding:"omitempty"`
}

type OfficialResponse struct {
	Torrents  []OfficialInfo `json:"torrents"`
	TotalPage int            `json:"total_page"`
}

// Official 获取官方种子文件列表 (GET /official)
func Official(c *gin.Context) {
	// 绑定参数
	var req OfficialRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	search := ""
	if req.Search != nil {
		search = *req.Search
	}

	// 获取种子列表
	bts, totalPage, err := db.GetTorrentList(
		db.TORRENT_ZONE_OFFICIAL, db.OrderByType(req.OrderBy), db.OrderType(req.Order),
		req.Page, req.PerPage, search)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, err.Error())
		log.Logger.Error("failed to get official torrent list" + err.Error())
		return
	}

	torrentsInfo := make([]OfficialInfo, 0, len(bts))
	for _, bt := range bts {
		torrentsInfo = append(torrentsInfo, OfficialInfo{
			CreatedAt: bt.CreatedAt.Format("2006-01-02 15:04:05"),
			Essay:     bt.Essay,
			Img:       bt.Img,
			Size:      util.ByteCountBinary(uint64(bt.Size)),
			Subtitle:  bt.Subtitle,
			Title:     bt.Title,
			TorrentID: bt.TorrentID,
			UpdateAt:  bt.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	resp.OKWithData(c, &OfficialResponse{
		Torrents:  torrentsInfo,
		TotalPage: totalPage,
	})
	log.Logger.Info("get official torrent list successfully")
}
