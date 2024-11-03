package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/go-the-complete-guide/11-course-project-rest-api/db"
	"github.com/paulfernandosr/go-the-complete-guide/11-course-project-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
