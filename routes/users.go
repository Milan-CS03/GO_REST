package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/Milan-CS03/GO_REST/models"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	err = user.Save()
	if err != nil {
		log.Fatalf("not saving %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "user": user})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login successful!", "user": user})

}
