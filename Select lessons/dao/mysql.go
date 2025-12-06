package dao

import (
	"Select_lessons/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/school?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Student{},
		&model.Lesson{},
		&model.Enrollment{},
		&model.User{})
	if err != nil {
		return err
	}
	return nil
}
