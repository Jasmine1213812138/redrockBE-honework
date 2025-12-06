package midware

import (
	"Select_lessons/model"
	"Select_lessons/respond"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	refreshSecret = os.Getenv("REFRESH_SECRET")
	accessSecret  = os.Getenv("ACCESS_SECRET")
)

func stripBearer(token string) string {
	if strings.HasPrefix(strings.ToLower(strings.TrimSpace(token)), "bearer") {
		return strings.TrimSpace(token[len("Bearer "):])
	}
	return strings.TrimSpace(token)
}
func VerifyTokens(tokenString string) (*model.CustomClaims, error) {
	raw := stripBearer(tokenString)
	token, err := jwt.ParseWithClaims(raw, &model.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return accessSecret, nil
	}, jwt.WithLeeway(time.Second*5))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.Type != "access" {
		return nil, errors.New("token type error")
	}
	return claims, nil
}
func VerifyRefreshTokens(tokenString string) (*model.CustomClaims, error) {
	raw := stripBearer(tokenString)
	token, err := jwt.ParseWithClaims(raw, &model.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return accessSecret, nil
	}, jwt.WithLeeway(time.Second*5))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.Type != "refresh" {
		return nil, errors.New("token type error")
	}
	return claims, nil
}
func VerifyAdminTokens(tokenString string) (*model.CustomClaims, error) {
	raw := stripBearer(tokenString)
	token, err := jwt.ParseWithClaims(raw, &model.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return accessSecret, nil
	}, jwt.WithLeeway(time.Second*5))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.Type != "access" {
		return nil, errors.New("token type error")
	}
	if claims.Role != "admin" {
		return nil, errors.New("token role error")
	}
	return claims, nil
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		_, err := VerifyTokens(tokenString)
		if err != nil {
			c.JSON(401, respond.HandleError(err, tokenString))
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthMiddleAdminWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		_, err := VerifyAdminTokens(tokenString)
		if err != nil {
			c.JSON(401, respond.HandleError(err, tokenString))
			c.Abort()
			return
		}
		c.Next()
	}
}
