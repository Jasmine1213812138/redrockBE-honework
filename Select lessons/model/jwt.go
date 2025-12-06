package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	Type     string `json:"type"`
	jwt.RegisteredClaims
}
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
