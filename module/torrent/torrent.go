package torrent

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/anacrolix/torrent/bencode"
)

// NewBitTorrentFilePath 通过文件路径创建 BitTorrentFile
func NewBitTorrentFilePath(torrentFilePath string) (*BitTorrentFile, error) {
	// io.Reader
	fileHeader, err := os.Open(torrentFilePath)
	if err != nil {
		return nil, err
	}
	return NewBitTorrentFile(fileHeader)
}

// NewBitTorrentFile 通过文件创建 BitTorrentFile
func NewBitTorrentFile(fileReader *os.File) (*BitTorrentFile, error) {
	decoder := bencode.NewDecoder(fileReader)

	bencodeTorrent := &BitTorrentFile{}
	decodeErr := decoder.Decode(bencodeTorrent)
	if decodeErr != nil {
		return nil, decodeErr
	}
	return bencodeTorrent, nil
}

// GetHash 获取 torrent 文件的 hash
func (bencodeTorrent *BitTorrentFile) GetHash() string {
	// marshal info part and calculate SHA1
	marshaledInfo, marshalErr := bencode.Marshal(bencodeTorrent.Info)
	if marshalErr != nil {
		return ""
	}
	return fmt.Sprintf("%x", sha1.Sum(marshaledInfo))
}

// RepackTorrent 重新打包 torrent 文件
func (bencodeTorrent *BitTorrentFile) RepackTorrent(editStrategy *BitTorrentFileEditStrategy) error {
	// Re-pack torrent
	if editStrategy.Comment != "" {
		bencodeTorrent.Comment = editStrategy.Comment
	}
	if editStrategy.InfoSource != "" {
		bencodeTorrent.Info.Source = editStrategy.InfoSource
	}

	bencodeTorrent.Info.Private = false

	return nil
}
