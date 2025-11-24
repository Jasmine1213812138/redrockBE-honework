package main

import (
	"fmt"
	"lesson04/list"
	"lesson04/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:070306and102817Wsl@tcp(127.0.0.1:3306)/school?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	fmt.Println("数据库连接成功")
	err = db.AutoMigrate(
		&models.Student{},
		&models.Lesson{},
		&models.Enrollment{})
	if err != nil {
		fmt.Println("建表失败")
	}
	fmt.Println("建表成功")
	db.Create(&models.Student{StudentId: 20252120, Name: "夜华"})
	db.Create(&models.Student{StudentId: 20252121, Name: "素素"})
	db.Create(&models.Student{StudentId: 20252122, Name: "糖宝"})
	db.Create(&models.Lesson{LessonId: 101, LessonName: "修仙", Capacity: 2, EnrolledCount: 0})
	db.Create(&models.Lesson{LessonId: 102, LessonName: "下凡历劫", Capacity: 1, EnrolledCount: 0})
	r := gin.Default()
	r.POST("/select", func(c *gin.Context) {
		studentId, _ := strconv.Atoi(c.Query("student_id"))
		lessonId, _ := strconv.Atoi(c.Query("lesson_id"))
		err := list.Select(studentId, lessonId, db)
		if err != nil {
			fmt.Println("出错了")
			c.JSON(400, gin.H{"message": "选课失败，请重新核对"})
			return
		}
		c.JSON(200, gin.H{"message": "选课成功"})
	})
	r.GET("/search", func(c *gin.Context) {
		var lessons []models.Lesson
		db.Find(&lessons)
		c.JSON(200, lessons)
	})
	r.Run(":8080")

}
