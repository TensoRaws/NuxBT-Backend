package torrent

import (
	"crypto/sha1"
	"fmt"
	"github.com/anacrolix/torrent/bencode"
	"io"
)

type BencodeTorrent struct {
	Announce  string `bencode:"announce"`
	CreatedBy string `bencode:"created by,omitempty"`
	CreatedAt int    `bencode:"creation date,omitempty"`
	Info      struct {
		Files []struct {
			Length uint64   `bencode:"length"`
			Path   []string `bencode:"path"`
		} `bencode:"files"`
		Name        string `bencode:"name"`
		Pieces      string `bencode:"pieces"`
		PieceLength uint64 `bencode:"piece length"`
		Private     int    `bencode:"private"`
		Source      string `bencode:"source"`
	} `bencode:"info"`
}

// RepackTorrent 重新打包 torrent 文件
func RepackTorrent(fileReader io.Reader) (*BencodeTorrent, string, error) {
	// Decode
	// See https://godoc.org/github.com/anacrolix/torrent/bencode
	decoder := bencode.NewDecoder(fileReader)
	bencodeTorrent := &BencodeTorrent{}
	decodeErr := decoder.Decode(bencodeTorrent)
	if decodeErr != nil {
		return nil, "", decodeErr
	}

	//// Re-pack torrent
	//// TODO: 根据配置修改
	bencodeTorrent.Announce = ""
	//bencodeTorrent.Info.Source = "[Alpha] SpiderX"
	//// 0: 公开种子 1: 私有种子
	//bencodeTorrent.Info.Private = 0

	// marshal info part and calculate SHA1
	marshaledInfo, marshalErr := bencode.Marshal(bencodeTorrent.Info)
	if marshalErr != nil {
		return nil, "", marshalErr
	}

	//metainfo.Hash
	bytesArray := sha1.Sum(marshaledInfo)
	hashString := fmt.Sprintf("%x", bytesArray[:])
	return bencodeTorrent, hashString, nil
}
