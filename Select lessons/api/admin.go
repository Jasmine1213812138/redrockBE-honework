package api

import (
	"Select_lessons/model"
	"Select_lessons/respond"
	"Select_lessons/sv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddLessons(c *gin.Context, db *gorm.DB) {
	var lessons model.Lesson
	if err := c.ShouldBindJSON(&lessons); err != nil {
		c.JSON(400, respond.HandleError(err, nil))
		return
	}
	err := sv.AddLesson(lessons, db)
	if err != nil {
		c.JSON(400, respond.HandleError(err, lessons))
	} else {
		c.JSON(200, respond.HandleError(err, lessons))
	}
}
func DeleteStudent(c *gin.Context, db *gorm.DB) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, respond.HandleError(err, nil))
		return
	}
	err := sv.AddStudent(student, db)
	if err != nil {
		c.JSON(400, respond.HandleError(err, student))
	} else {
		c.JSON(200, respond.HandleError(err, student))
	}
}
