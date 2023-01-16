package OneToOneTable

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID         uint
	Name       string
	Age        int
	Gender     bool
	UserInfoID uint     //外键
	UserInfo   UserInfo // 通过UserInfo可以拿到用户详情信息
}

type UserInfo struct {
	ID   uint
	User *User
	Addr string
	Like string
}

func CreateOneToOneTable(db *gorm.DB) {
	db.AutoMigrate(&UserInfo{}, &User{})
}

func AddDateOneToOneTable(db *gorm.DB) {
	// 1.添加用户，并自动添加用户详情
	db.Create(&User{
		Name:   "枫枫",
		Age:    21,
		Gender: true,
		UserInfo: UserInfo{
			Addr: "湖南省",
			Like: "写代码",
		},
	})
	//// 2.为已有用户添加或更新用户详情
	//var user User
	//db.Take(&user, 1)
	//db.Create(&UserInfo{
	//	User: &user, //必须要指定User，因此UserInfo结构体必须要包含一个User类字段(为了避免包含嵌套，用*User 地址类型)
	//	Addr: "北京市",
	//	Like: "睡觉",
	//})
}

func QueryOneToOneTable(db *gorm.DB) {
	// 查询用户，同时查询具体用户信息
	var user User
	db.Preload("UserInfo").Take(&user, 1)
	fmt.Println(user)

	// 查询信息，并查询所属用户
	var userInfo UserInfo
	db.Preload("User").Take(&userInfo, 2)
	data, _ := json.Marshal(userInfo)
	fmt.Println(string(data))
}

func DeleteOneToOneTable(db *gorm.DB) {
	//一对一表要删除，必须主表和副表同时删除
	var user User
	db.Take(&user)
	db.Debug().Select("UserInfo").Delete(&user)

	//var userInfo UserInfo
	//db.Take(&userInfo)
	//db.Debug().Select("User").Delete(&userInfo)
}
