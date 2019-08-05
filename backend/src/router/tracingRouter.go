package router

import (
	"controller"
	"github.com/gin-gonic/gin"
	"middleware"
)

func InitTracingRoutes(tracingGroup *gin.RouterGroup) {
	tracingGroup.POST("/",  middleware.IsAuth, controller.NewTracing)
}
