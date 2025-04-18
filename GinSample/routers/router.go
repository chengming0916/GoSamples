package routers

import (
	"GoSamples/GinSample/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// 静态页面
	r.StaticFile("/favicon.ico", "./assets/static/favicon.ico")
	r.Static("/assets", "./assets/static/assets")
	r.StaticFile("/", "./assets/static/index.html")

	// 用户路由表
	UserRouter(r)

	// 简单认证
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin", // user:password
	}))
	authorized.POST("user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"account":  "admin",
			"password": "admin",
		})
	})
}

var (
	userController = &controllers.UserController{}
)

func UserRouter(r *gin.Engine) {
	r.GET("/prod", userController.FindAll)
	r.POST("/prod", userController.Create)
}
