package models

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	Name	string `json:"name"`
	jwt.StandardClaims
}
