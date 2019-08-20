package router

import (
	"middleware"
	"controller"
	"github.com/gin-gonic/gin"
)

func InitLocationRoutes(locationGroup *gin.RouterGroup) {
	locationGroup.GET("/", middleware.IsAuth, controller.GetAllLocations)
	locationGroup.POST("/", middleware.IsAuth, controller.UpdateLocation)
}
