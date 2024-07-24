package torrent

import (
	"time"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/module/oss"
	"github.com/TensoRaws/NuxBT-Backend/module/torrent"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
)

type OfficialInfo struct {
	CreatedAt string `json:"created_at"`
	Essay     string `json:"essay"`
	Img       string `json:"img"`
	Size      string `json:"size"`
	Subtitle  string `json:"subtitle"`
	Title     string `json:"title"`
	TorrentID int32  `json:"torrent_id"`
	UpdateAt  string `json:"update_at"`
}

type Info struct {
	AnidbID     int32   `json:"anidb_id"`
	AudioCodec  string  `json:"audio_codec"`
	CreatedAt   string  `json:"created_at"`
	UpdateAt    string  `json:"update_at"`
	Description *string `json:"description,omitempty"`
	Essay       *string `json:"essay,omitempty"`
	Genre       string  `json:"genre"`
	Img         string  `json:"img"`
	Language    string  `json:"language"`
	Magnet      string  `json:"magnet"`
	Official    bool    `json:"official"`
	Resolution  string  `json:"resolution"`
	Size        string  `json:"size"`
	Status      string  `json:"status"`
	Subtitle    string  `json:"subtitle"`
	Title       string  `json:"title"`
	TorrentID   int32   `json:"torrent_id"`
	UploaderID  int32   `json:"uploader_id"`
	URL         *string `json:"url,omitempty"`
	VideoCodec  string  `json:"video_codec"`
}

// GetTorrentOSSKey 获取种子 OSS Key，参数为 hash
func GetTorrentOSSKey(hash string) string {
	return hash + ".torrent"
}

// GetTorrentOSSUrl 获取种子 OSS 地址
func GetTorrentOSSUrl(hash string, title string) (string, error) {
	ossUrl, err := oss.GetPresignedURL(GetTorrentOSSKey(hash), title+".torrent", 12*time.Hour)
	if err != nil {
		return "", err
	}

	return ossUrl, nil
}

// GetTorrentInfo 获取种子信息
func GetTorrentInfo(bt *model.Torrent) (*Info, error) {
	magnet := torrent.GetMagnet(bt.Hash, torrent.TRACKER_LIST)

	size := util.ByteCountBinary(uint64(bt.Size))

	urlString, err := GetTorrentOSSUrl(bt.Hash, bt.Title)
	if err != nil {
		return nil, err
	}

	return &Info{
		AnidbID:     bt.AnidbID,
		AudioCodec:  bt.AudioCodec,
		CreatedAt:   bt.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:    bt.UpdatedAt.Format("2006-01-02 15:04:05"),
		Description: &bt.Description,
		Essay:       &bt.Essay,
		Genre:       bt.Genre,
		Img:         bt.Img,
		Language:    bt.Language,
		Magnet:      magnet,
		Official:    bt.Official,
		Resolution:  bt.Resolution,
		Size:        size,
		Status:      bt.Status,
		Subtitle:    bt.Subtitle,
		Title:       bt.Title,
		TorrentID:   bt.TorrentID,
		UploaderID:  bt.UploaderID,
		URL:         &urlString,
		VideoCodec:  bt.VideoCodec,
	}, nil
}
