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
func GetTorrentInfo(t *model.Torrent) (*Info, error) {
	magnet := torrent.GetMagnet(t.Hash, torrent.TRACKER_LIST)

	size := util.ByteCountBinary(uint64(t.Size))

	urlString, err := GetTorrentOSSUrl(t.Hash, t.Title)
	if err != nil {
		return nil, err
	}

	return &Info{
		AnidbID:     t.AnidbID,
		AudioCodec:  t.AudioCodec,
		CreatedAt:   t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:    t.UpdatedAt.Format("2006-01-02 15:04:05"),
		Description: &t.Description,
		Essay:       &t.Essay,
		Genre:       t.Genre,
		Img:         t.Img,
		Language:    t.Language,
		Magnet:      magnet,
		Official:    t.Official,
		Resolution:  t.Resolution,
		Size:        size,
		Status:      t.Status,
		Subtitle:    t.Subtitle,
		Title:       t.Title,
		TorrentID:   t.TorrentID,
		UploaderID:  t.UploaderID,
		URL:         &urlString,
		VideoCodec:  t.VideoCodec,
	}, nil
}
