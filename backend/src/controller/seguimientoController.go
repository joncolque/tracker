package controller

import (
	"dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
)

func NewSeguimiento(c *gin.Context) {
	var seguimiento dto.SeguimientoDTO
	err := c.BindJSON(&seguimiento)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error1",//strErr(err),
		})
		return
	}
	err = service.CreateSeguimiento(&seguimiento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error2",//strErr(err),
		})
		return
	}
	seguimientoDTO, _ := service.FindSeguimientoByName(seguimiento.Name)
	c.JSON(http.StatusCreated, gin.H{
		"data": seguimientoDTO,
	})
}

func GetAllSeguimientos(c *gin.Context) {
	seguimientos, err := service.GetAllSeguimientos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error3",//strErr(err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": seguimientos,
	})
}
