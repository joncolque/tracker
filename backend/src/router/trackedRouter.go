package router

import (
	"controller"
	"github.com/gin-gonic/gin"
)

func InitTrackedRoutes(trackedGroup *gin.RouterGroup) {
	trackedGroup.POST("/", controller.NewTracked)
	trackedGroup.GET("/", controller.GetAllTracked)
}
