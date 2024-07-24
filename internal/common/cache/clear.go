package cache

import (
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"strconv"
)

// ClearCacheByKeys 清除缓存
func ClearCacheByKeys(keys ...string) {
	c := cache.Cache
	c.Del(keys...)
}

// ClearTorrentDetailCacheByTorrentID 清除种子详情缓存
func ClearTorrentDetailCacheByTorrentID(torrentID int32) {
	key := "/api/v1/torrent/detail?torrent_id=" + strconv.Itoa(int(torrentID))
	ClearCacheByKeys(key)
}
