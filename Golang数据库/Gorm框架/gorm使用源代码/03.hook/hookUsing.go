package hookUsing

import (
	"fmt"
	"gorm.io/gorm"
)

type StudentWithHook struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

// 在将结构体对象插入到数据库时，总会调用此方法
func (user *StudentWithHook) BeforeCreate(tx *gorm.DB) (err error) {
	email := fmt.Sprintf("%s@qq.com", user.Name)
	user.Email = &email
	return nil
}

func ShowHookEffect(DB *gorm.DB) {
	//根据对象建表
	s := new(StudentWithHook)
	DB.AutoMigrate(s)

	//插入一条数据
	email := "xxx@qq.com"
	student := StudentWithHook{
		Name:   "枫枫",
		Age:    21,
		Gender: true,
		Email:  &email,
	}
	DB.Create(&student)
}
