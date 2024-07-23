package torrent

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/anacrolix/torrent/bencode"
)

func NewBitTorrentFileFromMultipart(fh *multipart.FileHeader) (*BitTorrentFile, error) {
	fileReader, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer func(fileReader multipart.File) {
		err := fileReader.Close()
		if err != nil {
			fmt.Println("Failed to close multipart file")
		}
	}(fileReader)

	decoder := bencode.NewDecoder(fileReader)
	bencodeTorrent := &BitTorrentFile{}
	decodeErr := decoder.Decode(bencodeTorrent)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return bencodeTorrent, nil
}

// NewBitTorrentFileFromBytes 通过字节创建 BitTorrentFile
func NewBitTorrentFileFromBytes(torrentBytes []byte) (*BitTorrentFile, error) {
	torrentBytesReader := bytes.NewReader(torrentBytes)

	decoder := bencode.NewDecoder(torrentBytesReader)
	bencodeTorrent := &BitTorrentFile{}
	decodeErr := decoder.Decode(bencodeTorrent)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return bencodeTorrent, nil
}

// NewBitTorrentFileByPath 通过文件路径创建 BitTorrentFile
func NewBitTorrentFileByPath(torrentFilePath string) (*BitTorrentFile, error) {
	// io.Reader
	fileHeader, err := os.Open(torrentFilePath)
	defer func(fileHeader *os.File) {
		err := fileHeader.Close()
		if err != nil {
			fmt.Println("Failed to close file", fileHeader.Name())
		}
	}(fileHeader)

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

// Repack 重新打包 torrent 文件
func (bencodeTorrent *BitTorrentFile) Repack(editStrategy *BitTorrentFileEditStrategy) error {
	// Re-pack torrent
	if editStrategy.Announce != nil {
		bencodeTorrent.Announce = *editStrategy.Announce
	}
	if editStrategy.AnnounceList != nil {
		var announceList [][]string
		for _, announce := range editStrategy.AnnounceList {
			announceList = append(announceList, []string{announce})
		}
		bencodeTorrent.AnnounceList = announceList
	}
	if editStrategy.Comment != nil {
		bencodeTorrent.Comment = *editStrategy.Comment
	}
	if editStrategy.InfoSource != nil {
		bencodeTorrent.Info.Source = *editStrategy.InfoSource
	}

	bencodeTorrent.Info.Private = editStrategy.Private

	return nil
}

// GetFileList 获取 torrent 的文件列表和大小
func (bencodeTorrent *BitTorrentFile) GetFileList() []BitTorrentFileListItem {
	var fileList []BitTorrentFileListItem

	// 当 torrent 文件只有一个文件时，Info.Files 为空
	if len(bencodeTorrent.Info.Files) == 0 {
		fileList = append(fileList, BitTorrentFileListItem{
			Path: []string{bencodeTorrent.Info.Name},
			Size: util.ByteCountBinary(bencodeTorrent.Info.Length),
		})

		return fileList
	}

	// 当 torrent 文件有多个文件时，Info.Files 不为空，从中获取文件列表
	for _, file := range bencodeTorrent.Info.Files {
		fileList = append(fileList, BitTorrentFileListItem{
			Path: file.Path,
			Size: util.ByteCountBinary(file.Length),
		})
	}
	return fileList
}

// GetTotalSize 获取 torrent 文件的总大小，单位为字节
func (bencodeTorrent *BitTorrentFile) GetTotalSize() int64 {
	var totalSize uint64
	// 当 torrent 文件只有一个文件时，Info.Files 为空
	if len(bencodeTorrent.Info.Files) == 0 {
		totalSize = bencodeTorrent.Info.Length
		return int64(totalSize)
	}

	// 当 torrent 文件有多个文件时，Info.Files 不为空，从中获取文件列表
	for _, file := range bencodeTorrent.Info.Files {
		totalSize += file.Length
	}
	return int64(totalSize)
}

// ConvertToBytes 将 torrent 文件转换为字节
func (bencodeTorrent *BitTorrentFile) ConvertToBytes() ([]byte, error) {
	// Marshal the entire torrent file
	marshaledTorrent, marshalErr := bencode.Marshal(bencodeTorrent)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return marshaledTorrent, nil
}

// SaveTo 将 torrent 文件保存到指定路径
func (bencodeTorrent *BitTorrentFile) SaveTo(filePath string) error {
	// Marshal the entire torrent file
	marshaledTorrent, marshalErr := bencode.Marshal(bencodeTorrent)
	if marshalErr != nil {
		return marshalErr
	}

	// Write to file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Failed to close file", file.Name())
		}
	}(file)

	_, err = file.Write(marshaledTorrent)
	if err != nil {
		return err
	}

	return nil
}
