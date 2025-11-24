package models

// 学生表
type Student struct {
	StudentId int `gorm:"primaryKey"`
	Name      string
}

// 课程表
type Lesson struct {
	LessonId      int `gorm:"primaryKey"`
	LessonName    string
	Capacity      int
	EnrolledCount int
}

// 选课表
type Enrollment struct {
	Id        int `gorm:"primaryKey"`
	StudentId int `gorm:"foreignKey:StudentId"`
	LessonId  int `gorm:"foreignKey:LessonId"`
}
