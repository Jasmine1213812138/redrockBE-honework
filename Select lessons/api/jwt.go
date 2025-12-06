package api

import (
	"Select_lessons/midware"
	"Select_lessons/respond"
	"Select_lessons/utils"

	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	Token := c.GetHeader("Authorization")
	claims, err := midware.VerifyRefreshTokens(Token)
	if err != nil {
		c.JSON(400, "refresh token fail")
		return
	}
	var accessToken string
	accessToken, err = utils.GenerateToken(claims.UserName, claims.Role)
	if err != nil {
		c.JSON(400, respond.HandleError(err, "token fail"))
		return
	}
	c.JSON(200, respond.HandleError(err, accessToken))
}
