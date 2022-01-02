package model

import (
	"github.com/dgrijalva/jwt-go"
)

// User model
type User struct {
	ID    uint   `json:"id",gorm:"PrimaryKey"`
	Login string `json:"login"`
	// The password is stored as a hash number
	Password    uint32 `json:"password"`
	Permissions uint8
}

type CustomClaims struct {
	*jwt.StandardClaims

	ID    uint
	Login string
	Permissions uint8
}