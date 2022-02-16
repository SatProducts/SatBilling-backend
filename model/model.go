package model

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model

	Login       string `json:"login"`
	Password    string `json:"password"`
	Permissions uint8  `json:"permissions"`
	Vacation    bool   `json:"vacation"`
	Tasks       uint   `json:"tasks"`
}

type Claims struct {
	*jwt.StandardClaims

	ID          uint   `json:"id"`
	Login       string `json:"login"`
	Permissions uint8  `json:"permissions"`
}

type Task struct {
	gorm.Model

	ForUser  uint   `json:"for"`
	FromUser uint   `json:"from"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	Address  string `json:"address"`
}
