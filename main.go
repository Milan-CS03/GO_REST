package main

import (
	"github.com/Milan-CS03/GO_REST/db"
	"github.com/Milan-CS03/GO_REST/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
