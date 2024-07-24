package db

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
	"gorm.io/gorm"
)

// CheckTorrentExist 检查种子是否存在，确保 unique
func CheckTorrentExist(hash string) bool {
	q := query.Torrent
	_, err := q.Where(q.Hash.Eq(hash)).First()
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// CreateTorrent 创建种子
func CreateTorrent(torrent *model.Torrent) (err error) {
	q := query.Torrent
	err = q.Create(torrent)
	return err
}

// DeleteTorrent 删除种子
func DeleteTorrent(torrentID int32) (err error) {
	q := query.Torrent
	info, err := q.Where(q.TorrentID.Eq(torrentID)).Delete()
	if err != nil {
		return err
	}
	if info.RowsAffected == 0 {
		return fmt.Errorf("no rows affected, nothing will be updated, torrent ID: %v", torrentID)
	}
	return err
}

// PatchTorrent 更新种子信息，根据 torrentID 和 details 更新种子信息
func PatchTorrent[T *model.Torrent | map[string]any | any](torrentID int32, details T) (err error) {
	q := query.Torrent
	_, err = q.Where(q.TorrentID.Eq(torrentID)).Updates(details)
	if err != nil {
		return err
	}
	return nil
}

// GetTorrentByHash 根据 Hash 获取种子
func GetTorrentByHash(hash string) (*model.Torrent, error) {
	q := query.Torrent
	torrent, err := q.Where(q.Hash.Eq(hash)).First()
	return torrent, err
}

// GetTorrentByID 根据 ID 获取种子
func GetTorrentByID(torrentID int32) (*model.Torrent, error) {
	q := query.Torrent
	torrent, err := q.Where(q.TorrentID.Eq(torrentID)).First()
	return torrent, err
}

// GetTorrentList 获取种子列表，返回种子列表和分页数量
func GetTorrentList(zone TorrentZone, orderBy OrderByType, order OrderType,
	page int, perPage int, search string) ([]*model.Torrent, int, error) {
	if !order.Validate() || !orderBy.Validate() || !zone.Validate() {
		return nil, 0, fmt.Errorf("invalid order: %v, %v", order, orderBy)
	}

	q := query.Torrent
	var torrents []*model.Torrent
	var totalPages int

	// 分区
	var Query query.ITorrentDo
	switch zone {
	case TORRENT_ZONE_OFFICIAL:
		Query = q.Where(q.Official)
	case TORRENT_ZONE_GENERAL:
		Query = q.Where(q.Status.Eq(STATUS_APPROVED))
	case TORRENT_ZONE_PENDING:
		Query = q.Where(q.Status.Eq(STATUS_PENDING)).Or(q.Status.Eq(STATUS_REJECTED))
	default:
		return nil, 0, fmt.Errorf("invalid zone: %v", zone)
	}

	// 搜索
	var searchAnidbID int
	searchAnidbID, err := strconv.Atoi(search)
	if err != nil {
		searchAnidbID = 0
	}

	if search != "" {
		Query = Query.Where(q.Title.Like(search)).Or(q.Subtitle.Like(search)).
			Or(q.Hash.Eq(search)).Or(q.AnidbID.Eq(int32(searchAnidbID)))
	}

	// 获取总记录数
	count, err := Query.Count()
	if err != nil {
		return nil, 0, err
	}
	// 计算总页数
	totalPages = (int(count) + perPage - 1) / perPage
	offset := (page - 1) * perPage

	// 排序
	if orderBy == ORDER_BY_TYPE_DATE {
		if order == ORDER_TYPE_ASC {
			Query = Query.Order(q.CreatedAt.Asc())
		} else if order == ORDER_TYPE_DESC {
			Query = Query.Order(q.CreatedAt.Desc())
		}
	} else if orderBy == ORDER_BY_TYPE_SIZE {
		if order == ORDER_TYPE_ASC {
			Query = Query.Order(q.Size.Asc())
		} else if order == ORDER_TYPE_DESC {
			Query = Query.Order(q.Size.Desc())
		}
	}

	torrents, err = Query.Limit(perPage).Offset(offset).Find()
	if err != nil {
		return nil, 0, err
	}

	return torrents, totalPages, nil
}
