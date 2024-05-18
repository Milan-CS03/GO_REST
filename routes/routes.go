package routes

import (
	"github.com/Milan-CS03/GO_REST/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middleware.Authenticate, createEvents)
	server.PUT("/events/:id", middleware.Authenticate, updateEvent)
	server.DELETE("events/:id", middleware.Authenticate, deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.POST("/events/:id/register", middleware.Authenticate, registerEvents)
	server.DELETE("/events/:id/register", middleware.Authenticate, cancelRegistration)

}
