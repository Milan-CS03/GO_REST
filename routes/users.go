package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Milan-CS03/GO_REST/models"
	"github.com/Milan-CS03/GO_REST/utils"
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
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user"})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful!", "token": token})

}
