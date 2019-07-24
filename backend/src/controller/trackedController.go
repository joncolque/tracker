package controller

import (
	"dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
)

func NewTracked(c *gin.Context) {
	var tracked dto.TrackedDTO
	err := c.BindJSON(&tracked)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error1",//strErr(err),
		})
		return
	}
	err = service.CreateTracked(&tracked)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error2",//strErr(err),
		})
		return
	}
	trackedDTO, _ := service.FindTrackedById(tracked.Id_tracked)
	c.JSON(http.StatusCreated, gin.H{
		"data": trackedDTO,
	})
}

func GetAllTracked(c *gin.Context) {
	trackeds, err := service.GetAllTracked()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error3",//strErr(err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": trackeds,
	})
}
