package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vilkovtato/rest-api/db"
	"github.com/vilkovtato/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
