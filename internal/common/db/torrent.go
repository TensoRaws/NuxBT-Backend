package db

import (
	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
)

// CreateTorrent 创建种子
func CreateTorrent(torrent *model.Torrent) (err error) {
	q := query.Torrent
	err = q.Create(torrent)
	return err
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
