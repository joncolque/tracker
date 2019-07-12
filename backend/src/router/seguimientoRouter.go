package router

import (
	"controller"
	"github.com/gin-gonic/gin"
)

func InitSeguimientoRoutes(seguimientoGroup *gin.RouterGroup) {
	seguimientoGroup.POST("/", controller.NewSeguimiento)
	seguimientoGroup.GET("/", controller.GetAllSeguimientos)
}
