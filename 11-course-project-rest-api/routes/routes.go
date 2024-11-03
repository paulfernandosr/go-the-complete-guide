package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paulfernandosr/go-the-complete-guide/11-course-project-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/events")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/", createEvent)
	authenticated.PUT("/:id", updateEvent)
	authenticated.DELETE("/:id", deleteEvent)
	authenticated.POST("/:id/register", registerForEvent)
	authenticated.DELETE("/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
