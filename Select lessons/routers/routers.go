package routers

import (
	"Select_lessons/api"
	"Select_lessons/midware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	setupAuth(r, db)
	setupCourse(r, db)
	setupAdmin(r, db)
	return r
}
func setupAuth(r *gin.Engine, db *gorm.DB) {
	r.POST("/login", func(c *gin.Context) {
		api.UserLogin(c, db)
	})
	r.POST("/register", func(c *gin.Context) {
		api.UserRegister(c, db)
	})
	r.POST("/refresh", func(c *gin.Context) {
		api.RefreshToken(c)
	})
}
func setupCourse(r *gin.Engine, db *gorm.DB) {
	courseGroup := r.Group("/course")
	courseGroup.Use(midware.AuthMiddleware())
	{
		courseGroup.GET("all", func(c *gin.Context) {
			api.GetLessons(c, db)
		})
		courseGroup.GET("my course", func(c *gin.Context) {
			api.EnrolledLessons(c, db)
		})
		courseGroup.POST("/select", func(c *gin.Context) {
			api.SelectLessons(c, db)
		})
		courseGroup.POST("/delete", func(c *gin.Context) {
			api.DeleteLessons(c, db)
		})
	}
}
func setupAdmin(r *gin.Engine, db *gorm.DB) {
	adminGroup := r.Group("/admin")
	adminGroup.Use(midware.AuthMiddleAdminWare())
	{
		adminGroup.POST("addLessons", func(c *gin.Context) {
			api.AddLessons(c, db)
		})
		adminGroup.DELETE("deleteStudent", func(c *gin.Context) {
			api.DeleteStudent(c, db)
		})
	}

}
