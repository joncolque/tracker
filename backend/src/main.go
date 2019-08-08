package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"config"
	"router"
	"dbconnection"
	"deliverer"
)

func main(){
	config.InitConfig()
	dbconnection.InitMongoDB()
	defer dbconnection.CloseMongoDB()
	initAPI()
}

func initAPI() {
	fmt.Println("INIT API")
	app := gin.Default()
	
	app.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin, token",
		ExposedHeaders: "",
		Credentials: true,
		ValidateHeaders: false,
	}))

	router.InitDeviceRoutes(app.Group("/device"))
	router.InitTracingRoutes(app.Group("/tracing"))
	router.InitPositionRoutes(app.Group("/position"))

	deliverer.InitProcess()

	app.Run(":8080")
}