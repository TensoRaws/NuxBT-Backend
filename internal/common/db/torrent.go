package db

import (
	"errors"

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
