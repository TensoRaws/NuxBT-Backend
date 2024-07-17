package torrent

import (
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"os"
	"testing"
)

func TestRepackTorrent(t *testing.T) {
	torrentFilePath := "test.torrent"

	// io.Reader
	fileHeader, err := os.Open(torrentFilePath)
	torrent, s, err := RepackTorrent(fileHeader)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(s)
	s = util.StructToString(torrent)
	t.Log(s)
}
