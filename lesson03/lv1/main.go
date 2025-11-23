package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/talk", func(c *gin.Context) {
		receive := c.Query("msg")
		var response string
		if receive == "ping" {
			response = "pong"
		} else if receive == "hellosever" {
			response = "hellloclient"
		} else {
			response = ""
		}
		c.JSON(200, gin.H{
			"data": response,
		})
	})
	err := r.Run()
	if err != nil {
		fmt.Println("启动失败")
	}
}
