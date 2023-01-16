package singleTable

import "gorm.io/gorm"

func SingleSaveUpdateAll(DB *gorm.DB) {
	var stu Student
	DB.Take(&stu)
	stu.Name = "无敌枫枫"
	email := "12345@qq.com"
	stu.Email = &email
	stu.Age = 0        //即使是类型默认的零值也会被更新
	stu.Gender = false //即使是类型默认的零值也会被更新
	DB.Save(&stu)      //一次性更新stu的所有字段
}

func SingleSaveUpdate(DB *gorm.DB) {
	var stu Student
	DB.Take(&stu)
	stu.Name = "万年老二枫枫"
	stu.Age = 110
	stu.Gender = true
	DB.Select("name").Save(&stu)
}

func BatchUpdateRow(DB *gorm.DB) {
	var stuList []Student
	DB.Find(&stuList, "age >= ?", 25).Update("email", "xxx@163.com") //批量更新多行数据的某一列

	var stu Student
	DB.Model(&stu).Where("age >= ?", 25).Update("email", "xxx@google.com") //效果与上述语言相同
}

func BatchUpdateRowsByStruct(DB *gorm.DB) {
	//如果是结构体，它默认不会更新零值
	var stuList []Student
	email := "xxx@163.com"
	DB.Find(&stuList, "age >= ?", 25).Updates(Student{
		Email:  &email,
		Gender: false, // bool类型的默认值，因此不会更新
	})

	//如果想让他更新零值，需要结合使用select
	email = "xxx@google.com"
	DB.Find(&stuList, "age >= ?", 25).Select("email", "gender").Updates(Student{
		Email:  &email,
		Gender: true, // 被select强制指定了，即使是类型默认值也会被更新
	})
}

func BatchUpdateRowsByMap(DB *gorm.DB) {
	//如果是map，它默认会更新零值
	var stuList []Student
	email := "xxx@163.com"
	DB.Find(&stuList, "age >= ?", 25).Updates(map[string]any{
		"email":  &email,
		"gender": false,
	})
}
