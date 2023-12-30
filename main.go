package main

import (
	"web-service/connections"
	"web-service/routers"

	"github.com/gin-gonic/gin"

	Middlewares "web-service/utils/middlewares"
)

func main() {
	connections.ConnectDB()
	if connections.ConnectionError != nil {
		panic(connections.ConnectionError)
	} // connections.ConnectionError

	server := gin.Default()

	server.Use(Middlewares.Log())

	v1 := server.Group("/api/v1")
	{
		routers.UserRoutes(v1)
		routers.StoreRoutes(v1)
		routers.AuthRoutes(v1)
	}

	server.Run(":9091")
}
