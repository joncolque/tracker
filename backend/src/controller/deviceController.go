package controller

import (
	"dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
)

func NewDevice(c *gin.Context) {

	var device dto.DeviceDTO
	var completeDevice dto.CompleteDeviceDTO

	id_user := c.Param("id_user")

	//TODO: Tomar la "app" del JWT_Auth del Header
	app:= "flowtrace"

	err := c.BindJSON(&device)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	completeDevice.Id_device = device.Id_device
	completeDevice.Id_user = id_user
	completeDevice.App = app

	err = service.CreateDevice(completeDevice)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "the device could not be registered",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
			"message": "the device has been registered successfully",
		})
	return
}

func GetAllDevices(c *gin.Context) {
	devices, err := service.GetAllDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error3",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": devices,
	})
}
