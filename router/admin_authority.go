package router

import (
	"FuckingVersion1/api/v1"
	"FuckingVersion1/middleware"
	"github.com/gin-gonic/gin"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("authority").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		AuthorityRouter.POST("createAuthority", v1.CreateAuthority)   //创建角色
		AuthorityRouter.POST("deleteAuthority", v1.DeleteAuthority)   //删除角色
		AuthorityRouter.POST("getAuthorityList", v1.GetAuthorityList) //获取角色列表
		AuthorityRouter.POST("setDataAuthority", v1.SetDataAuthority) //设置角色资源权限
	}
}
