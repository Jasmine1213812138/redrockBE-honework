package dao

import (
	"Select_lessons/model"
	"Select_lessons/respond"
	"errors"

	"gorm.io/gorm"
)

func CreateStudent(student model.Student, db *gorm.DB) error {
	err := db.Create(&student).Error
	if err != nil {
		return err
	}
	return nil
}
func CreateUser(user model.User, db *gorm.DB) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func SearchStudent(id uint64, db *gorm.DB) (*model.Student, error) {
	var student model.Student
	err := db.Where("student_id=?", id).First(&student).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, respond.ResourceNotFound
		}
		return nil, err
	}
	return &student, nil
}
func SearchUser(username string, db *gorm.DB) (*model.User, error) {
	var user model.User
	err := db.Where("user_name=?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, respond.UserNotFound
		}
		return nil, err
	}
	return &user, nil
}
