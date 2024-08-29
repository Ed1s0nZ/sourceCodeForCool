package routes

import (
	"coolv0.1/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) {
	r.Static("/statics", "./statics")
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/favicon.ico", "./statics/favicon.ico")
	r.GET("/", controller.IndexHandler)
	r.POST("/", controller.Time1, controller.BypassHandler())

}
