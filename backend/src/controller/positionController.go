package controller

import (
	"dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
	// "fmt"
)

func UpdatePosition(c *gin.Context) {
	// Del JWT saco que usuario es y su aplicación. c->Header->Authorization
	id_user := "user_1"
	app := "flowtrace"

	var sLocation dto.SignatureLocationDTO
	err := c.BindJSON(&sLocation)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	var userLocation dto.UserLocationDTO
	userLocation.Id_user = id_user
	userLocation.App = app
	userLocation.Location.Longitude = sLocation.Longitude
	userLocation.Location.Latitude = sLocation.Latitude
	userLocation.Location.Timestamp = sLocation.Timestamp

	service.InsertLocation(userLocation)

	c.JSON(http.StatusOK, gin.H{
		"message": "the location has been updated successfully",
	})
}

func GetAllLocations(c *gin.Context) {

	var sFilterLocation dto.SignatureFilterLocationDTO
	err := c.Bind(&sFilterLocation)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	var filterLocation dto.FilterLocationDTO
	filterLocation.Id_user = sFilterLocation.Id_user

	locations, err := service.GetAllLocations(filterLocation)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Can not get the data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": locations,
	})

}
