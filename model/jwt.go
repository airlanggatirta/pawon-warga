package model

import (
	"github.com/dgrijalva/jwt-go"
)

type KitabisaClaims struct {
	User User `json:"userinfo"`
	jwt.StandardClaims
}

type NewKitabisaClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}
