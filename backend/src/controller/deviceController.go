package controller

import (
	"dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
)

func NewDevice(c *gin.Context) {

	var sDevice dto.SignatureDeviceDTO
	id_user := c.Param("id_user")

	//TODO: Tomar la "app" del JWT_Auth c->Header->Authorization

	app:= "flowtrace"

	err := c.BindJSON(&sDevice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	var uDevice dto.UpsertDeviceDTO
	uDevice.Id_device = sDevice.Id_device
	uDevice.Id_user = id_user
	uDevice.App = app

	err = service.CreateDevice(uDevice)

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
