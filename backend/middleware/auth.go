package middleware

import (
	"better_calendar_backend/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := strings.Split(authHeader, " ")[1]
	if authHeader != "" {
		claims := jwt.MapClaims{}
		user, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("helloworld"), nil
		})
		utils.ErrorHandler(err)
		c.Set("userId", int(user.Claims.(jwt.MapClaims)["userId"].(float64)))
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
}
