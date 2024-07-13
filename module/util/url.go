package util

import (
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"net/url"
)

// RemoveQueryParameter 从 URL 中移除一组 query 参数
func RemoveQueryParameter(rawurl string, keys ...string) string {
	u, err := url.Parse(rawurl)
	if err != nil {
		// 处理错误
		log.Logger.Error("failed to parse url: " + err.Error())
		return rawurl
	}

	// 获取原始的 query 参数
	query := u.Query()

	// 删除 query 参数
	for _, key := range keys {
		query.Del(key)
	}

	// 重新设置 URL 的 query 参数
	u.RawQuery = query.Encode()

	return u.String()
}
