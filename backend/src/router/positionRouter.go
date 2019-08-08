package router

import (
	"middleware"
	"controller"
	"github.com/gin-gonic/gin"
)

func InitPositionRoutes(positionGroup *gin.RouterGroup) {
	positionGroup.GET("/", middleware.IsAuth, controller.GetAllLocations)
	positionGroup.POST("/", middleware.IsAuth, controller.UpdatePosition)
}
