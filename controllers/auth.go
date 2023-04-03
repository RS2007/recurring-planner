package controllers

import (
	// "better_calendar_backend/db"

	"better_calendar_backend/db"
	"better_calendar_backend/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func GoogleAuth(c *gin.Context) {
	code := c.Query("code")
	Db := db.GetDB()
	fmt.Println("body:", c.Request.Header)
	fmt.Println("code", code)
	tokenRes, err := utils.GetGoogleAuthToken(code)
	utils.ErrorHandler(err)
	userFromGoogle, err := utils.GetGoogleUser(tokenRes.AccessToken, tokenRes.IdToken)
	utils.ErrorHandler(err)
	fmt.Println("user: ", userFromGoogle)
	queryRow := Db.QueryRow("SELECT * FROM users WHERE email=$1;", userFromGoogle.Email)
	var user db.UserModelWithId
	err = queryRow.Scan(&user.UserId, &user.Name, &user.Email, &user.AccentColor)

	if err != nil {
		if err == sql.ErrNoRows {
			user = db.UserModelWithId{
				UserId: 1,
				UserModel: db.UserModel{
					Name:        userFromGoogle.Name,
					Email:       userFromGoogle.Email,
					AccentColor: "blue",
				},
			}

			var userId int32
			err := Db.QueryRow("INSERT INTO users (name,email,accentColor) VALUES ($1,$2,$3) RETURNING userId;", userFromGoogle.Name, userFromGoogle.Email, "blue").Scan(&userId)
			utils.ErrorHandler(err)
			user.UserId = int32(userId)
			utils.ErrorHandler(err)
		} else {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
	}

	claims := jwt.MapClaims{
		"userId": user.UserId,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	err = godotenv.Load()
	utils.ErrorHandler(err)
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := jwtToken.SignedString([]byte(jwtSecret))
	utils.ErrorHandler(err)
	jsonResponse, err := json.Marshal(gin.H{"token": tokenString, "name": user.Name, "accentColor": user.AccentColor, "email": user.Email, "googleAccessToken": tokenRes.AccessToken, "googleIdToken": tokenRes.IdToken})
	utils.ErrorHandler(err)
	c.SetCookie("RP_USER_DETAILS", string(jsonResponse), 3600, "/", "https://recurring-planner-monorepo.onrender.com/", false, false)
	c.Redirect(http.StatusPermanentRedirect, "/")
}
