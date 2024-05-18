package middleware

import (
	"net/http"

	"github.com/Milan-CS03/GO_REST/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}
	uid, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}
	context.Set("uid", uid)
	context.Next()

}
