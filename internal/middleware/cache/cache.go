package cache

import (
	"bytes"
	"sync"

	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/gin-gonic/gin"
)

var once sync.Once

var Clients = map[cache.RDB]*cache.Client{
	cache.IPLimit: {},
	cache.User:    {},
}

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	// 向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	// 完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func Init() {
	once.Do(func() {
		cache.NewRedisClients(Clients)
	})
}
