package routes

import (
	"net/http"
	"strconv"

	"github.com/Milan-CS03/GO_REST/models"
	"github.com/gin-gonic/gin"
)

func registerEvents(context *gin.Context) {
	userID := context.GetInt64("uid")

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse ID"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Find event with ID"})
		return
	}
	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "registration successful"})

}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("uid")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	var event models.Event
	event.ID = eventID
	err = event.CancelRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "cancelled"})

}
