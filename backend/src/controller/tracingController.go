package controller

import (
	"dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
	// "fmt"
)

func NewTracing(c *gin.Context) {
	var tracingDevice dto.TrackedDTO
	err := c.BindJSON(&tracingDevice)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error1",//strErr(err),
		})
		return
	}

	service.InitTracing(tracingDevice.Id_tracked)

}
