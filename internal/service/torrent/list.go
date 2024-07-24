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

type ListRequest struct {
	Order   string  `form:"order" binding:"required,oneof=asc desc"`
	OrderBy string  `form:"order_by" binding:"required,oneof=date size"`
	Page    int     `form:"page" binding:"required,min=1"`
	PerPage int     `form:"per_page" binding:"required,min=10"`
	Search  *string `form:"search" binding:"omitempty"`
	Zone    string  `form:"zone" binding:"required,oneof=official general pending"`
}

type ListResponse struct {
	Torrents  []Info `json:"torrents"`
	TotalPage int    `json:"total_page"`
}

// List 获取种子文件列表 (GET /list)
func List(c *gin.Context) {
	// 绑定参数
	var req ListRequest
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
		db.TorrentZone(req.Zone), db.OrderByType(req.OrderBy), db.OrderType(req.Order),
		req.Page, req.PerPage, search)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, err.Error())
		log.Logger.Error("failed to get torrent list" + err.Error())
		return
	}

	torrentsInfo := make([]Info, 0, len(bts))
	for _, bt := range bts {
		torrentsInfo = append(torrentsInfo, Info{
			AnidbID:    bt.AnidbID,
			AudioCodec: bt.AudioCodec,
			CreatedAt:  bt.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:   bt.UpdatedAt.Format("2006-01-02 15:04:05"),
			Genre:      bt.Genre,
			Img:        bt.Img,
			Language:   bt.Language,
			Magnet:     torrent.GetMagnet(bt.Hash, torrent.TRACKER_LIST),
			Official:   bt.Official,
			Resolution: bt.Resolution,
			Size:       util.ByteCountBinary(uint64(bt.Size)),
			Status:     bt.Status,
			Subtitle:   bt.Subtitle,
			Title:      bt.Title,
			TorrentID:  bt.TorrentID,
			UploaderID: bt.UploaderID,
			VideoCodec: bt.VideoCodec,
		})
	}

	resp.OKWithData(c, &ListResponse{
		Torrents:  torrentsInfo,
		TotalPage: totalPage,
	})
	log.Logger.Info("get torrent list successfully")
}
