package MultiTableQuery

import (
	"fmt"
	"gorm.io/gorm"
)

func QueryOneToMul(db *gorm.DB) {

	// 常规查询
	var user User
	db.Take(&user) //这样的查询方式无法查询到外键对应的另一个表的信息
	fmt.Println(user)

	// 预加载，使用预加载的方式来加载用户列表
	user = User{}
	db.Preload("Articles").Take(&user, 1) //预加载的字符串参数就是外键所关联的字段名
	fmt.Println(user)

	// 预加载，使用预加载的方式来加载文章
	var article Article
	db.Preload("User").Take(&article, 1)
	fmt.Println(article)

	// 嵌套预加载，预加载的字符串参数可以多次嵌套
	// 查询文章，显示用户，并且显示用户关联的所有文章
	article = Article{}
	db.Preload("User.Articles").Take(&article, 1)
	fmt.Println(article)

	// 条件预加载，对外键管理的另一表的数据进行过滤
	user = User{}
	db.Preload("Articles", "id = ?", 1).Take(&user, 1) //仅获取 id = 1 的文章
	fmt.Println(user)

	// 使用匿名函数进行自定义预加载
	user = User{}
	db.Preload("Articles", func(db *gorm.DB) *gorm.DB {
		return db.Where("id in ?", []int{1, 2}) //仅获取 id = 1 or id = 2 的文章
	}).Take(&user, 1)
	fmt.Println(user)
}

func DeleteOneToMul(db *gorm.DB) {

	// 级联删除，删除用户，与用户关联的文章也一并删除
	var user User
	db.Take(&user, 1)
	db.Select("Articles").Delete(&user)

	// 清除外键关系，删除用户，与将与用户关联的文章，外键设置为null
	user = User{}
	db.Preload("Articles").Take(&user, 2)
	db.Model(&user).Association("Articles").Delete(&user.Articles)
	db.Delete(&user)
}
