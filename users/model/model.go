package model

import (
	"gorm.io/gorm"
	"github.com/dgrijalva/jwt-go"
)

// User model
type User struct {
	gorm.Model

	ID    uint   `json:"id",gorm:"PrimaryKey"`
	Login string `json:"login"`
	// The password is stored as a hash number
	Password string `json:"password"`
	
	Permissions uint8  `json:"permissions"`
	Vacation    bool   `json:"vacation"`
}

type CustomClaims struct {
	*jwt.StandardClaims
	
	ID    uint
	Login string
	Permissions uint8
}