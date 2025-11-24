package list

import (
	"errors"
	"lesson04/models"

	"gorm.io/gorm"
)

func Select(studentId int, lessonId int, db *gorm.DB) error {
	//把选课系统改造为函数来方便接口引用
	/*db, err := gorm.Open(mysql.Open("root:070306and102817Wsl@tcp(127.0.0.1:3306)/school?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
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
	var student models.Student
	var lesson models.Lesson
	此为观察是否创建成功
		db.Find(&students)
		for _, student := range students {
			fmt.Println(student)
	}

	夜华选了下凡历劫
	db.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("student_id=?", "20252120").First(&student)
		result1 := tx.Where("lesson_id=?", "102").First(&lesson)
		if result.Error != nil {
			return errors.New("不存在该学生，请重新检查")
		}
		if result1.Error != nil {
			return errors.New("不存在课程，请重新检查")
		}
		if lesson.Capacity <= lesson.EnrolledCount {
			return errors.New("抱歉，已经选满了")
		}
		tx.Create(&models.Enrollment{StudentId: 20252120, LessonId: 102})
		tx.Model(&models.Lesson{}).
			Where("lesson_id=?", "102").
			Update("enroll_count", gorm.Expr("enrolled_count + ?", 1))
		return nil
	})
	*/
	var student models.Student
	var lesson models.Lesson
	return db.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("student_id=?", "studentId").First(&student)
		result1 := tx.Where("lesson_id=?", "lessonId").First(&lesson)
		if result.Error != nil {
			return errors.New("不存在该学生，请重新检查")
		}
		if result1.Error != nil {
			return errors.New("不存在课程，请重新检查")
		}
		if lesson.Capacity <= lesson.EnrolledCount {
			return errors.New("抱歉，已经选满了")
		}
		tx.Create(&models.Enrollment{StudentId: studentId, LessonId: lessonId})
		tx.Model(&models.Lesson{}).
			Where("lesson_id=?", lessonId).
			Update("enrolled_count", gorm.Expr("enrolled_count + ?", 1))
		return nil
	})

}
