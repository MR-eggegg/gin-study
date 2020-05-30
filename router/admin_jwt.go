package router

import (
	"FuckingVersion1/api/v1"
	"FuckingVersion1/middleware"
	"github.com/gin-gonic/gin"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) //jwt加入黑名单
	}
}
