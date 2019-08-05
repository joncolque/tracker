package controller

import (
	"dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
	"fmt"
)

func NewTracing(c *gin.Context) {
	var tracingDevices dto.TracingDevicesDTO
	err := c.BindJSON(&tracingDevices)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	if tracingDevices.Follow {
		usersOnTrack, err := service.Tracing(tracingDevices)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "The devices could not be registered",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": usersOnTrack,
		})

	} else {
		users, err := service.StopTracing(tracingDevices)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "The devices could not be unregistered",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})

	}
	
	return

}
