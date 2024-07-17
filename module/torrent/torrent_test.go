package torrent

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, s, "7f3956e5a15b34b62159727c08f944c7e433ad1e")
	assert.Equal(t, torrent.Info.Name, "lenna.jpg")
}
