package torrent

import "mime/multipart"

type UploadRequest struct {
	File        multipart.FileHeader `form:"torrent_file" type:"blob" binding:"required"`
	AnidbID     int32                `form:"anidb_id" binding:"required"`
	AudioCodec  string               `form:"audio_codec" binding:"required,oneof FLAC AAC AC3 DTS DDP LPCM other"`
	Description string               `form:"description" binding:"required"`
	Essay       *string              `form:"essay" binding:"omitempty"`
	Genre       string               `form:"genre" binding:"required,oneof BDrip WEBrip DVDrip Remux Blu-ray WEB-DL DVD HDTV other"`
	Img         string               `form:"img" binding:"required"`
	Language    string               `form:"language" binding:"required,oneof Chinese English Japanese other"`
	Official    bool                 `form:"official" binding:"required"`
	Resolution  string               `form:"resolution" binding:"required,oneof 480p 720p 1080p 2160p other"`
	Subtitle    string               `form:"subtitle" binding:"required"`
	Title       string               `form:"title" binding:"required"`
	VideoCodec  string               `form:"video_codec" binding:"required,oneof H.265 H.264 AV1 VP9 other"`
}
