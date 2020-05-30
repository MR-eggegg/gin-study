package initialize

import (
	"FuckingVersion1/global"
	"FuckingVersion1/middleware"
	"FuckingVersion1/router"
	"github.com/gin-gonic/gin"

)

func Routers() *gin.Engine{
	var Routers = gin.Default()

	global.MyLog.Debug("use middleware logger")
	// 跨域
	Routers.Use(middleware.Cors())
	global.MyLog.Debug("use middleware cors")
	ApiGroup :=Routers.Group("")
	{
		router.InitBaseRouter(ApiGroup)
		router.InitUserRouter(ApiGroup)
		router.InitJwtRouter(ApiGroup)
		router.InitAuthorityRouter(ApiGroup)
		global.MyLog.Info("router register success")
	}
	return Routers
}