package router

import (
	v1 "FuckingVersion1/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("ping2", v1.Ping2)
		BaseRouter.POST("register", v1.Register)
		BaseRouter.POST("login", v1.Login)
		//BaseRouter.POST("captcha", v1.Captcha)
		//BaseRouter.GET("captcha/:captchaId", v1.CaptchaImg)
	}
	return BaseRouter
}

