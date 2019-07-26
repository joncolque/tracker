package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func IsAuth(c *gin.Context) {
	//TODO: Validar JWT_Auth del header con el servicio de Auth.
	fmt.Println("Valido JWT_Auth")
	c.Next()
}
