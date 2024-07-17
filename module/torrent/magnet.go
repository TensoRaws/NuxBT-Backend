package torrent

import (
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/types/infohash"
)

// GetMagnet returns a magnet link for the given torrent hash and trackers.
func GetMagnet(torrentHash string, trackers []string) string {
	mag := metainfo.Magnet{
		InfoHash: infohash.FromHexString(torrentHash),
		Trackers: trackers,
	}

	return mag.String()
}
