package torrent

import (
	"crypto/sha1"
	"fmt"
	"github.com/anacrolix/torrent/bencode"
	"io"
)

type TorrentFile struct {
	Announce     string   `bencode:"announce"`
	AnnounceList []string `bencode:"announce-list,omitempty"`
	CreationDate int64    `bencode:"creation date,omitempty"`
	Comment      string   `bencode:"comment,omitempty"`
	CreatedBy    string   `bencode:"created by,omitempty"`
	Info         struct {
		Files []struct {
			Path   []string `bencode:"path"`
			Length uint64   `bencode:"length"`
		} `bencode:"files,omitempty"`
		Name        string `bencode:"name"`
		Pieces      []byte `bencode:"pieces"`
		PieceLength int64  `bencode:"piece length"`
		Length      uint64 `bencode:"length,omitempty"`
		Private     bool   `bencode:"private,omitempty"`
		Source      string `bencode:"source,omitempty"`
	} `bencode:"info"`
}

// RepackTorrent 重新打包 torrent 文件
func RepackTorrent(fileReader io.Reader) (*TorrentFile, string, error) {
	// Decode
	// See https://godoc.org/github.com/anacrolix/torrent/bencode
	decoder := bencode.NewDecoder(fileReader)

	bencodeTorrent := &TorrentFile{}
	decodeErr := decoder.Decode(bencodeTorrent)
	if decodeErr != nil {
		return nil, "", decodeErr
	}

	//// Re-pack torrent
	//// TODO: 根据配置修改
	//bencodeTorrent.Announce = ""
	//bencodeTorrent.Info.Source = "[Alpha] SpiderX"
	//// 0: 公开种子 1: 私有种子
	//bencodeTorrent.Info.Private = 0

	// marshal info part and calculate SHA1
	marshaledInfo, marshalErr := bencode.Marshal(bencodeTorrent.Info)
	if marshalErr != nil {
		return nil, "", marshalErr
	}

	return bencodeTorrent, fmt.Sprintf("%x", sha1.Sum(marshaledInfo)), nil
}