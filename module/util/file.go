package util

import (
	"github.com/dustin/go-humanize"
)

// ByteCountBinary 把 Length 转换为文件大小, 自适应单位
func ByteCountBinary(b uint64) string {
	return humanize.Bytes(b)
}
