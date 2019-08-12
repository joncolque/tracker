package router

import (
	"middleware"
	"controller"
	"github.com/gin-gonic/gin"
)

func InitLocationRoutes(positionGroup *gin.RouterGroup) {
	positionGroup.GET("/", middleware.IsAuth, controller.GetAllLocations)
	positionGroup.POST("/", middleware.IsAuth, controller.UpdatePosition)
}
