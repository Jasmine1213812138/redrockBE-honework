package dao

import (
	"Select_lessons/model"
	"Select_lessons/respond"
	"errors"
	"sync"

	"gorm.io/gorm"
)

func AddLesson(db *gorm.DB, lesson model.Lesson) error {
	err := db.Create(&lesson)
	if err != nil {
		return err.Error
	}
	return nil
}
func GetLesson(db *gorm.DB) ([]model.Lesson, error) {
	var lessons []model.Lesson
	err := db.Find(&lessons).Error
	if err != nil {
		return nil, err
	}
	return lessons, nil
}
func GetEnrollment(db *gorm.DB, studentId int) ([]int, error) {
	var enrolledLessons []int
	err := db.Model(&model.Enrollment{}).
		Where("student_id = ?", studentId).
		Pluck("lesson_id", &enrolledLessons).Error
	if err != nil {
		return nil, err
	}
	return enrolledLessons, nil
}
func EnrolledLessons(db *gorm.DB, lessonId []int) ([]model.Lesson, error) {
	var lessons []model.Lesson
	err := db.Where("lesson_id IN (?)", lessonId).Find(&lessons).Error
	if err != nil {
		return nil, err
	}
	return lessons, nil
}
func SearchLessons(db *gorm.DB, lessonId uint64) (*model.Lesson, error) {
	err := db.Where("lesson_id = ?", lessonId).Find(&model.Lesson{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, respond.ResourceNotFound
		}
		return nil, err
	}
	return &model.Lesson{}, err
}
func Select(db *gorm.DB, lesson *model.Lesson, enrollment *model.Enrollment, mu *sync.Mutex) error {
	mu.Lock()
	defer mu.Unlock()
	if lesson.Capacity > lesson.EnrolledCount {
		err := db.Create(&enrollment).Error
		if err != nil {
			return err
		}
		err = db.Model(lesson).Where("lesson_id=?", enrollment.LessonId).
			Update("enrolled_count", gorm.Expr("enrolled_count+1")).Error
		if err != nil {
			return err
		}
	} else {
		return respond.NoLeft
	}
	return nil
}
func Delete(db *gorm.DB, lesson *model.Lesson, enrollment *model.Enrollment, mu *sync.Mutex) error {
	mu.Lock()
	defer mu.Unlock()
	err := db.Where("id=?", enrollment.Id).Delete(&model.Enrollment{}).Error

	if err != nil {
		return err
	}
	err = db.Model(lesson).Where("lesson_id=?", enrollment.LessonId).
		Update("enrolled_count", gorm.Expr("enrolled_count-1")).Error
	if err != nil {
		return err
	}
	return nil
}
