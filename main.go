package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Milan-CS03/GO_REST/db"
	"github.com/Milan-CS03/GO_REST/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)

	server.Run(":8080")
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No such  event with given ID present"})
		return
	}
	events, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, *events)

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	//event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		log.Fatalf("not saving %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create events111"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})

}
