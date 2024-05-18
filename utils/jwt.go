package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "abcSecretKey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretkey))
}

func VerifyToken(token string) (int64, error) {
	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected sign in method")
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}
	if !parsedtoken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claim")

	}
	//email := claims["email"].(string)
	userID := int64(claims["userID"].(float64))

	return userID, nil

}
