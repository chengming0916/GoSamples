package routers

import (
	"GoSamples/GinSample/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// 用户路由表
	UserRouter(r)
}

var (
	userController = &controllers.UserController{}
)

func UserRouter(r *gin.Engine) {
	r.GET("/prod", userController.FindAll)
	r.POST("/prod", userController.Create)
}
