package role

const (
	// 管理角色
	ADMIN = "ADMIN" // 管理员，拥有访问管理的权限
	// 普通角色
	VIP           = "VIP"           // VIP，有一些可能有用的权限
	ADVANCED_USER = "ADVANCED_USER" // 高级用户，有一些有用的权限
	// 特殊角色
	UPLOADER = "UPLOADER" // 发种员，可以直接发种
	REVIEWER = "REVIEWER" // 审种员，可以审核种子
)
