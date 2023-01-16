package advancedQuery

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	singleTable "gorm_exec/02.singleTable"
)

func PtrString(str string) *string {
	return &str
}

func CreateTable(db *gorm.DB) {
	var stuList []singleTable.Student
	db.Find(&stuList).Delete(&stuList) //删除历史数据,被删除数据的表名由结构体名称确定(Student --> students)

	stuList = []singleTable.Student{
		{ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com"), Gender: true},
		{ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn"), Gender: true},
		{ID: 3, Name: "枫枫", Age: 23, Email: PtrString("ff@yahoo.com"), Gender: true},
		{ID: 4, Name: "刘大", Age: 54, Email: PtrString("liuda@qq.com"), Gender: true},
		{ID: 5, Name: "李武", Age: 23, Email: PtrString("liwu@lly.cn"), Gender: true},
		{ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn"), Gender: false},
		{ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com"), Gender: false},
		{ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com"), Gender: false},
		{ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com"), Gender: true},
	}
	db.Create(&stuList)
}

func AdvancedQuery(db *gorm.DB) {
	var stuList []singleTable.Student
	db.Where("name = ?", "枫枫").Find(&stuList)
	fmt.Println(stuList)

	count := db.Not("name = ?", "枫枫").Find(&stuList).RowsAffected
	fmt.Println("count = ", count, stuList)

	db.Where("name in (?)", []string{"如燕", "李元芳"}).Find(&stuList)
	fmt.Println(stuList)

	db.Where("name like ?", "李%").Find(&stuList)
	fmt.Println(stuList)

	//db.Where("age > ? and email like ?", 23, "%@qq.com").Find(&stuList)
	db.Where("age > ?", 23).Where("email like ?", "%@qq.com").Find(&stuList)
	fmt.Println(stuList)

	//db.Where("gender = ? or email like ?", false, "%@qq.com").Find(&stuList)
	db.Where("gender = ?", false).Or("email like ?", "%@qq.com").Find(&stuList)
	fmt.Println(stuList)
}

func QueryWithStructOrMap(db *gorm.DB) {
	var stuList []singleTable.Student

	db.Debug().Where(&singleTable.Student{Name: "李元芳", Age: 0}).Find(&stuList)
	fmt.Println(stuList)

	db.Debug().Where(map[string]any{
		"name": "李元芳",
		"age":  0,
	}).Find(&stuList)
	fmt.Println(stuList)
}

func QueryWithSelect(db *gorm.DB) {
	var stuList []singleTable.Student

	db.Debug().Select("name", "age").Find(&stuList)
	fmt.Println(stuList)

	type User struct {
		Name string
		Age  int
	}
	var userList []User
	//db.Debug().Select("name", "age").Find(&stuList).Scan(&userList)
	//fmt.Println("改进前", userList)

	userList = make([]User, 0)
	db.Debug().Model(&singleTable.Student{}).Select("name", "age").Scan(&userList)
	fmt.Println("使用Model()进行改进", userList)

	userList = make([]User, 0)
	db.Debug().Table("students").Select("name", "age").Scan(&userList)
	fmt.Println("使用Table()进行改进", userList)

}

func OrderAndPageAndDistinct(db *gorm.DB) {

	// 给数据排序
	var stuList []singleTable.Student
	db.Order("age desc").Find(&stuList) //按照年龄进行desc降序排序 asc是升序
	fmt.Println(stuList)

	// 分页查询
	// 对应的SQL语句为： select * from students limit x offset y  或者 select * from students limit x,y
	stuList = make([]singleTable.Student, 0)

	db.Limit(2).Offset(0).Find(&stuList) //查询第一页的两条数据
	fmt.Println(stuList)

	db.Limit(2).Offset(2).Find(&stuList) //查询第二页的两条数据
	fmt.Println(stuList)

	// 更加通用的分页写法
	stuList = make([]singleTable.Student, 0)

	limit := 2
	page := 1
	offset := (page - 1) * limit

	db.Limit(limit).Offset(offset).Find(&stuList)
	fmt.Println(stuList)

	// 去重
	ageList := make([]int, 0)
	db.Table("students").Select("age").Distinct("age").Find(&ageList)
	fmt.Println(ageList)

	ageList = make([]int, 0)
	db.Table("students").Select("distinct age").Find(&ageList)
	fmt.Println(ageList)
}

func GroupQuery(db *gorm.DB) {

	// 按照性别进行简单分组查询各自人数
	var ageList []int
	db.Table("students").Select("count(id)").Group("gender").Scan(&ageList)
	fmt.Println(ageList)

	// 附带性别
	type AgeGroup struct {
		Count  int
		Gender int
	}
	ageGroup := []AgeGroup{}
	db.Debug().Table("students").Select("count(id) as count,?", "gender").
		Group("gender").
		Scan(&ageGroup)
	fmt.Println(ageGroup)

	// 进一步精确，附带所有男女生的name
	type NAgeGroup struct {
		Count  int
		Gender int
		Name   string `gorm:"column:group_concat(name)"`
	}
	NageGroup := []NAgeGroup{}
	db.Debug().Table("students").Select("count(id) as count", "gender",
		"group_concat(name)").Group("gender").Scan(&NageGroup)
	fmt.Println(NageGroup)

	NageGroup = []NAgeGroup{}
	db.Table("students").
		Raw("SELECT count(id) as count,gender,group_concat(name) FROM `students` GROUP BY `gender`").
		Scan(&NageGroup)
	fmt.Println(NageGroup)
}

func Subquery(DB *gorm.DB) {
	var stuList []singleTable.Student
	// 使用原生SQL进行子查询
	DB.Raw("select * from students where age > (select avg(age) from students)").Find(&stuList)
	fmt.Println(stuList)

	// 使用gorm进行子查询
	stuList = make([]singleTable.Student, 0)
	DB.Model(&singleTable.Student{}).Where("age > (?)", DB.Model(&singleTable.Student{}).
		Select("avg(age)")).Find(&stuList)
	fmt.Println(stuList)
}

func NamedParams(DB *gorm.DB) {

	var stuList []singleTable.Student
	// 传统方式：使用 ? 代表参数。缺点是参数不是很明确，不直观
	DB.Where("name = ? and age = ?", "枫枫", 23).Find(&stuList)
	fmt.Println(stuList)

	var mapList = map[string]any{}
	DB.Table("students").Where("name = ? and age = ?", "枫枫", 23).Find(&mapList)
	fmt.Println(mapList)

	// 使用命名参数的方式 + sql.Named()方法
	stuList = make([]singleTable.Student, 0)
	DB.Where("name = @name and age = @age", sql.Named("name", "枫枫"),
		sql.Named("age", 23)).Find(&stuList)
	fmt.Println(stuList)

	// 使用命名参数的方式 + map
	stuList = make([]singleTable.Student, 0)
	DB.Where("name = @name and age = @age",
		map[string]any{
			"name": "枫枫",
			"age":  23,
		}).Find(&stuList)
	fmt.Println(stuList)
}

func CallChaining(DB *gorm.DB) {
	var stuList []singleTable.Student
	DB.Where("name = ? and age = ?", "枫枫", 23).Find(&stuList)
	fmt.Println(stuList)

	// 链式调用
	DB.Scopes(ageEqual23).Scopes(nameEqualFF).Find(&stuList)
	DB.Scopes(ageEqual23, nameEqualFF).Find(&stuList)
	fmt.Println(stuList)
}

func ageEqual23(DB *gorm.DB) *gorm.DB {
	return DB.Where(" age = ?", 23)
}

func nameEqualFF(DB *gorm.DB) *gorm.DB {
	return DB.Where("name = ?", "枫枫")
}
