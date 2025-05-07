package main

import (
	"gin_rest_api/db"
	"gin_rest_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":9090")
}
