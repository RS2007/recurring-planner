package db

type UserModel struct {
	Name        string `json:"name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	AccentColor string `json:"accentColor" binding:"required"`
}
