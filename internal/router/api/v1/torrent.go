package v1

import (
	"time"

	middleware_cache "github.com/TensoRaws/NuxBT-Backend/internal/middleware/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/rbac"
	torrent_service "github.com/TensoRaws/NuxBT-Backend/internal/service/torrent"
	"github.com/TensoRaws/NuxBT-Backend/module/role"
	"github.com/gin-gonic/gin"
)

// TorrentRouterGroup 种子路由组
func TorrentRouterGroup(api *gin.RouterGroup) {
	torrent := api.Group("torrent/")

	// 上传种子
	torrent.POST("upload",
		jwt.RequireAuth(false),
		rbac.RABC(role.ADMIN, role.UPLOADER, role.ADVANCED_USER),
		torrent_service.Upload)

	// 获取种子文件列表
	torrent.GET("filelist",
		jwt.RequireAuth(false),
		middleware_cache.Response(1*time.Hour),
		torrent_service.FileList)
}
