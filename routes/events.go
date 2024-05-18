package routes

import (
	"github.com/Milan-CS03/GO_REST/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse ID"})
		return
	}
	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		log.Fatalf("not updating %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event udpated", "event": updatedEvent})

}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No such event with given ID present"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, *event)

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
