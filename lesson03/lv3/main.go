package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Studentgrade struct {
	Name  string    `json:"name"`
	Score []float64 `json:"grade"`
}
type Average struct {
	Average float64 `json:"average"`
}
type Errorjudge struct {
	Error string `json:"error"`
}

func main() {
	r := gin.Default()
	r.POST("/grade", func(c *gin.Context) {
		var student Studentgrade
		err := c.BindJSON(&student)
		if err != nil {
			log.Println("绑定失败")
			c.JSON(400, Errorjudge{
				Error: "抱歉，请按照格式输入",
			})
			return
		}
		if len(student.Score) == 0 {
			c.JSON(400, Errorjudge{
				Error: "成绩列表不能为空",
			})
			return
		}

		var sum float64
		for _, score := range student.Score {
			sum += score
		}
		average := sum / float64(len(student.Score))
		result := Average{
			Average: average}
		c.JSON(200, result)
	})
	err := r.Run(":8080")
	if err != nil {
		log.Println("启动失败")
	}
}
