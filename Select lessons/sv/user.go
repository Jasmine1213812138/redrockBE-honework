package sv

import (
	"Select_lessons/dao"
	"Select_lessons/model"
	"Select_lessons/utils"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

func UserRegister(user model.User, db *gorm.DB) error {
	var err error
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	err = dao.CreateUser(user, db)
	if err != nil {
		return err
	}
	return nil
}
func UserLogin(user model.User, db *gorm.DB) (error, bool) {
	findUser, err := dao.SearchUser(user.Username, db)
	if err != nil {
		return err, false
	}
	var result bool
	result, err = utils.CheckPasswordHash(user.Password, findUser.Password)
	return err, result
}
func TaskChan(enrollment *[]model.Enrollment) chan *model.Enrollment {
	ch := make(chan *model.Enrollment, 100)
	go func() {
		for _, e := range *enrollment {
			ch <- &e
		}
		close(ch)
	}()
	return ch
}
func GetLessons(db *gorm.DB) ([]model.Lesson, error) {
	lessons, err := dao.GetLesson(db)
	if err != nil {
		return nil, err
	}
	return lessons, nil
}
func EnrolledLessons(db *gorm.DB, studentId int) ([]model.Lesson, error) {
	id, err := dao.GetEnrollment(db, studentId)
	if err != nil {
		return nil, err
	}
	var result []model.Lesson
	result, err = dao.EnrolledLessons(db, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func SelectLessons(db *gorm.DB, ch chan *model.Enrollment) error {
	mu := &sync.Mutex{}
	resultChan := make(chan error, 100)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for task := range ch {
				_, errStudent := dao.SearchStudent(task.StudentId, db)
				lesson, errLesson := dao.SearchLessons(db, task.LessonId)
				errSelect := dao.Select(db, lesson, task, mu)
				if errStudent != nil {
					resultChan <- errStudent
				}
				if errLesson != nil {
					resultChan <- errLesson
				}
				if errSelect != nil {
					resultChan <- errSelect
				}
			}
		}()
	}
	wg.Wait()
	close(resultChan)
	var errors []error
	for err := range resultChan {
		errors = append(errors, err)
	}
	if len(errors) > 0 {
		return fmt.Errorf("选课失败数：%d", len(errors))
	}
	return nil
}
func BackLessons(db *gorm.DB, ch chan *model.Enrollment) error {
	mu := &sync.Mutex{}
	resultChan := make(chan error, 100)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for task := range ch {
				_, errStudent := dao.SearchStudent(task.StudentId, db)
				lesson, errLesson := dao.SearchLessons(db, task.LessonId)
				errBack := dao.Delete(db, lesson, task, mu)
				if errStudent != nil {
					resultChan <- errStudent
				}
				if errLesson != nil {
					resultChan <- errLesson
				}
				if errBack != nil {
					resultChan <- errBack
				}
			}
		}()
	}
	wg.Wait()
	close(resultChan)
	var errors []error
	if len(errors) > 0 {
		return fmt.Errorf("退课失败数：%d", len(errors))
	}
	return nil
}
