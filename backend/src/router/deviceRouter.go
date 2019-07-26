package router

import (
	"middleware"
	"controller"
	"github.com/gin-gonic/gin"
)

func InitDeviceRoutes(deviceGroup *gin.RouterGroup) {
	deviceGroup.PUT("/:id_user", middleware.IsAuth, controller.NewDevice)
	deviceGroup.GET("/", controller.GetAllDevices)
}
