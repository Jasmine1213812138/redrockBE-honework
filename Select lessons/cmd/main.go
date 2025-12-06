package main

import (
	"Select_lessons/model"
	"Select_lessons/routers"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/school?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(
		&model.Student{},
		&model.Lesson{},
		&model.Enrollment{},
		&model.User{})
	if err != nil {
		log.Fatal(err)
	}
	r := routers.SetUpRouter(db)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
