package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/catjpg", func(c *gin.Context) {
		file, err := os.ReadFile("cat.jpg")
		if err != nil {
			log.Println("未找到数据")
		}
		c.Header("Content-Type", "image/jpeg")
		c.Data(200, "image/jpeg", file)
	})
	err1 := r.Run(":8080")
	if err1 != nil {
		log.Println("启动失败")
	}
}

