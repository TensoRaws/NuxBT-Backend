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
	// 种子编辑
	torrent.POST("edit",
		jwt.RequireAuth(false),
		rbac.RABC(),
		torrent_service.Edit)
	// 种子删除
	torrent.POST("delete",
		jwt.RequireAuth(false),
		rbac.RABC(role.ADMIN),
		torrent_service.Delete)
	// 种子审核
	torrent.POST("review",
		jwt.RequireAuth(false),
		rbac.RABC(role.REVIEWER),
		torrent_service.Review)
	// 获取种子文件列表
	torrent.GET("filelist",
		jwt.RequireAuth(false),
		middleware_cache.Response(24*time.Hour),
		torrent_service.FileList)
}
