package middleware

import (
	"better_calendar_backend/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	err := godotenv.Load()
	utils.ErrorHandler(err)
	jwtSecretString := os.Getenv("JWT_SECRET")
	fmt.Println(authHeader)
	authHeaderSplit := strings.Split(authHeader, " ")
	if len(authHeaderSplit) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	token := authHeaderSplit[1]
	if authHeader != "" {
		claims := jwt.MapClaims{}
		user, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretString), nil
		})
		utils.ErrorHandler(err)
		c.Set("userId", int(user.Claims.(jwt.MapClaims)["userId"].(float64)))
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
}
