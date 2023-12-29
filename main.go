package main

import (
	"web-service/connections"
	"web-service/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	connections.ConnectDB()
	if connections.ConnectionError != nil {
		panic(connections.ConnectionError)
	} // connections.ConnectionError

	server := gin.Default()

	v1 := server.Group("/api/v1")
	{
		routers.UserRoutes(v1)
	}

	server.Run(":9091")
}
