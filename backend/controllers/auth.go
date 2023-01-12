package controllers

import (
	// "better_calendar_backend/db"
	"better_calendar_backend/db"
	"better_calendar_backend/utils"
	"database/sql"
	"fmt"
	"net/http"

	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func Register(c *gin.Context) {
	var newUser db.UserModel
	db := db.GetDB()
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Malformed request body"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	utils.ErrorHandler(err)
	stmt, err := db.Prepare(fmt.Sprintf("INSERT INTO users (name,password,email,accentColor) VALUES ('%s','%s','%s','%s');", newUser.Name, string(hashedPassword), newUser.Email, newUser.AccentColor))
	utils.ErrorHandler(err)
	result, err := stmt.Exec()
	fmt.Println(result)
	utils.ErrorHandler(err)
	c.JSON(http.StatusOK, gin.H{"message": "Registration is successful"})
}

func Login(c *gin.Context) {
	err := godotenv.Load()
	jwtSecret := os.Getenv("JWT_SECRET")
	utils.ErrorHandler(err)
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "malformed input"})
		return
	}
	Db := db.GetDB()
	var user struct {
		db.UserModel
		UserId int32
		ApiKey sql.NullString
	}
	queryRow := Db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE email='%s';", credentials.Email))
	err = queryRow.Scan(&user.UserId, &user.Name, &user.Password, &user.Email, &user.AccentColor, &user.ApiKey)
	fmt.Println(user)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "This email is not associated with a valid account"})
			return
		} else {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect username or password"})
		return
	}
	claims := jwt.MapClaims{
		"userId": user.UserId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	utils.ErrorHandler(err)
	c.IndentedJSON(http.StatusOK, gin.H{"token": tokenString, "name": user.Name, "accentColor": user.AccentColor})
}
