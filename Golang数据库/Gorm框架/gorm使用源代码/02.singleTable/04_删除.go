package singleTable

import "gorm.io/gorm"

func DeleteWithQueryResult(DB *gorm.DB) {
	//删除单条
	var stu Student
	DB.Take(&stu)
	DB.Delete(&stu)

	//批量删除
	var stuList []Student
	DB.Find(&stuList, []int{11, 12})
	DB.Delete(&stuList)
}

func DeleteByCondition(DB *gorm.DB) {
	DB.Delete(&Student{}, []int{6, 7})
}
