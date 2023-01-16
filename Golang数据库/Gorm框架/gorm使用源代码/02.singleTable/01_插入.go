package singleTable

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func CreateTable(DB *gorm.DB) {
	s := new(Student)
	DB.AutoMigrate(s)
}

func InsertSingle(DB *gorm.DB) {
	email := "xxx@qq.com"
	// 创建记录
	student := Student{
		Name:   "枫枫",
		Age:    21,
		Gender: true,
		Email:  &email,
	}
	DB.Create(&student)
}

func BatchInsert(DB *gorm.DB) {
	email := "xxx@qq.com"

	var stuList []Student

	for i := 0; i < 10; i++ {
		stuList = append(stuList, Student{
			Name:   fmt.Sprintf("学生%d", i+1),
			Age:    20 + i,
			Gender: true,
			Email:  &email,
		})
	}
	DB.Create(stuList)

}
