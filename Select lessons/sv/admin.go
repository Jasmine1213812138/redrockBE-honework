package sv

import (
	"Select_lessons/dao"
	"Select_lessons/model"

	"gorm.io/gorm"
)

func AddLesson(lesson model.Lesson, db *gorm.DB) error {
	err := dao.AddLesson(db, lesson)
	if err != nil {
		return err
	}
	return nil
}
func AddStudent(student model.Student, db *gorm.DB) error {
	err := dao.CreateStudent(student, db)
	if err != nil {
		return err
	}
	return nil
}
