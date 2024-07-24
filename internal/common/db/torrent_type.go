package db

const (
	STATUS_PENDING  = "pending"
	STATUS_APPROVED = "approved"
	STATUS_REJECTED = "rejected"
)

// OrderType 用于排序的升降序类型
type OrderType string

const (
	ORDER_TYPE_ASC  OrderType = "asc"
	ORDER_TYPE_DESC OrderType = "desc"
)

func (o OrderType) Validate() bool {
	return o == "asc" || o == "desc"
}

// OrderByType 用于排序的字段类型
type OrderByType string

const (
	ORDER_BY_TYPE_DATE OrderByType = "date"
	ORDER_BY_TYPE_SIZE OrderByType = "size"
)

func (o OrderByType) Validate() bool {
	return o == "date" || o == "size"
}

// TorrentZone 种子区域类型
type TorrentZone string

const (
	TORRENT_ZONE_OFFICIAL TorrentZone = "official"
	TORRENT_ZONE_GENERAL  TorrentZone = "general"
	TORRENT_ZONE_PENDING  TorrentZone = "pending"
)

func (t TorrentZone) Validate() bool {
	return t == "official" || t == "general" || t == "pending"
}
