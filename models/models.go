package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Title   string `json:"title"`
	Address string `json:"address"`
	UserID  uint   `json:"userID"`
	User    User   `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	gorm.Model

	Name string `json:"name"`
}
