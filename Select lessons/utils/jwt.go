package utils

import (
	"Select_lessons/model"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessSecret  = os.Getenv("ACCESS_SECRET")
	refreshSecret = os.Getenv("REFRESH_SECRET")
	issuer        = "Select lessons"
	accessTTL     = 15 * time.Minute
	refreshTTL    = 15 * 24 * time.Hour
)

func GenerateToken(userName string, role string) (accessToken string, err error) {
	now := time.Now()
	accessClaims := model.CustomClaims{
		UserName: userName,
		Role:     role,
		Type:     "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   fmt.Sprintf("%s", userName),
			Audience:  []string{"user"},
			ExpiresAt: jwt.NewNumericDate(now.Add(accessTTL)),
			NotBefore: jwt.NewNumericDate(now.Add(-5 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	accessTok := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessTok.SignedString(accessSecret)
	if err != nil {
		return "", fmt.Errorf("sign access token fail")
	}
	return accessToken, nil
}
func RefreshToken(userName string, role string) (refreshToken string, err error) {
	now := time.Now()
	refreshClaims := model.CustomClaims{
		UserName: userName,
		Role:     role,
		Type:     "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   fmt.Sprintf("%s", userName),
			Audience:  []string{"user"},
			ExpiresAt: jwt.NewNumericDate(now.Add(refreshTTL)),
			NotBefore: jwt.NewNumericDate(now.Add(-5 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	refreshTok := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshTok.SignedString(refreshSecret)
	if err != nil {
		return "", fmt.Errorf("sign refresh token fail")
	}
	return refreshToken, nil
}
