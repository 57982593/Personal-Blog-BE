package router

import (
	"github.com/gin-gonic/gin"
	"goTestProject/Personal-Blog/controllers"
)

func Router(e *gin.Engine)  {
	control := controllers.Controller{}
	api := e.Group("/api")

	api.GET("/downloadMain", control.DownloadMainExe)
	api.POST("/login", control.Login)
	api.POST("/register", control.Register)
}