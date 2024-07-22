package torrent

type BitTorrentFile struct {
	Announce     string     `bencode:"announce"`
	AnnounceList [][]string `bencode:"announce-list,omitempty"`
	CreationDate int64      `bencode:"creation date,omitempty"`
	Comment      string     `bencode:"comment,omitempty"`
	CreatedBy    string     `bencode:"created by,omitempty"`
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

type BitTorrentFileList struct {
	Path []string `json:"path"`
	Size string   `json:"size"`
}

type BitTorrentFileEditStrategy struct {
	Announce     *string
	AnnounceList []string
	Comment      *string
	Private      bool
	InfoSource   *string
}
