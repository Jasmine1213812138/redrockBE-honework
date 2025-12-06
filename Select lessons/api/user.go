package api

import (
	"Select_lessons/model"
	"Select_lessons/respond"
	"Select_lessons/sv"
	"Select_lessons/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRegister(c *gin.Context, db *gorm.DB) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, respond.HandleError(respond.Response{Status: "400", Info: "信息格式错误"}, nil))
		return
	}
	err := sv.UserRegister(user, db)
	if err != nil {
		c.JSON(400, respond.HandleError(err, nil))
		return
	}
	c.JSON(200, respond.HandleError(err, user.Username))

}
func UserLogin(c *gin.Context, db *gorm.DB) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, respond.HandleError(respond.Response{Status: "400", Info: "信息格式错误"}, nil))
		return
	}
	err, data := sv.UserLogin(user, db)
	if err != nil {
		c.JSON(500, respond.HandleError(err, data))
		return
	}
	var accessToken string
	var refreshToken string
	accessToken, err = utils.GenerateToken(user.Username, user.Role)
	if err != nil {
		c.JSON(500, respond.HandleError(err, "accessToken生成失败"))
		return
	}
	refreshToken, err = utils.RefreshToken(user.Username, user.Role)
	if err != nil {
		c.JSON(500, respond.HandleError(err, "refreshToken生成失败"))
		return
	}
	Token := model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(200, respond.HandleError(err, Token))

}

func GetLessons(c *gin.Context, db *gorm.DB) {
	lessons, err := sv.GetLessons(db)
	if err != nil {
		c.JSON(400, respond.HandleError(err, lessons))
	} else {
		c.JSON(200, respond.HandleError(err, lessons))
	}

}
func EnrolledLessons(c *gin.Context, db *gorm.DB) {
	studentId := c.Query("student_id")
	Id, err := strconv.Atoi(studentId)
	lessons, err := sv.EnrolledLessons(db, Id)
	if err != nil {
		c.JSON(400, respond.HandleError(err, lessons))
	} else {
		c.JSON(200, respond.HandleError(err, lessons))
	}
}
func SelectLessons(c *gin.Context, db *gorm.DB) {
	var tasks []model.Enrollment
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(400, respond.HandleError(err, nil))
		return
	}
	ch := sv.TaskChan(&tasks)
	err := sv.SelectLessons(db, ch)
	if err != nil {
		c.JSON(400, respond.HandleError(err, tasks))
	} else {
		c.JSON(200, respond.HandleError(err, tasks))
	}
}
func DeleteLessons(c *gin.Context, db *gorm.DB) {
	var tasks []model.Enrollment
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(400, respond.HandleError(err, nil))
		return
	}
	ch := sv.TaskChan(&tasks)
	err := sv.BackLessons(db, ch)
	if err != nil {
		c.JSON(400, respond.HandleError(err, tasks))
	} else {
		c.JSON(200, respond.HandleError(err, tasks))
	}
}
