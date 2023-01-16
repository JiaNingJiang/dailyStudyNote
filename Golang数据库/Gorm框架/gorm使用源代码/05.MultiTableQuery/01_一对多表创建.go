package MultiTableQuery

import "gorm.io/gorm"

type User struct {
	ID       uint      `gorm:"size:4"`
	Name     string    `gorm:"size:8"`
	Articles []Article // 用户拥有的文章列表,Article的切片,因此确定User
}

type Article struct {
	ID     uint   `gorm:"size:4"`
	Title  string `gorm:"size:16"`
	UserID uint   `gorm:"size:4"` // 属于   这里的类型要和引用的外键类型一致，包括大小
	User   User   // 属于
}

//type User struct {
//	ID       uint      `gorm:"size:4"`
//	Name     string    `gorm:"size:8;index"`
//	Articles []Article `gorm:"foreignKey:UserName;references:Name"` // 用户拥有的文章列表
//}
//
//type Article struct {
//	ID       uint   `gorm:"size:4"`
//	Title    string `gorm:"size:16"`
//	UserName string `gorm:"size:8"`
//	User     User   `gorm:"foreignKey:UserName;references:Name"` // 属于
//}

func CreateOneToMul(DB *gorm.DB) {
	DB.AutoMigrate(&User{}, &Article{})
}
