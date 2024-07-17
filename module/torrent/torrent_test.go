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
	torrent, sha1, err := RepackTorrent(fileHeader)
	if err != nil {
		return
	}
	t.Log(sha1)

	// 测试数据：种子文件的 SHA1 值
	//expectedSHA1 := "611b8701b68790a3f22a9e27bd1e9a047f78691b"
	//assert.Equal(t, expectedSHA1, sha1)
	s := util.StructToString(torrent)
	t.Log(s)
}
