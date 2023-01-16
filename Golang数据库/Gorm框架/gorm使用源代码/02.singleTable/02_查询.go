package singleTable

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func QuerySingle(DB *gorm.DB) {
	var stu1 Student

	DB.Take(&stu1)
	fmt.Println(stu1)
}

func QuerySingleWithLog(DB *gorm.DB) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // （日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 使用彩色打印
		},
	)
	session := DB.Session(&gorm.Session{Logger: newLogger})

	var stu Student
	session.Take(&stu)
	fmt.Println("Session Take:", stu) //获取一条数据

	stu = Student{}
	session.First(&stu)
	fmt.Println("Session First:", stu) //获取第一条数据(根据主键排序)

	stu = Student{}
	session.Last(&stu)
	fmt.Println("Session Last:", stu) //获取最后一条数据
}

func ConditionalQuery(DB *gorm.DB) {
	var stu Student
	DB.Take(&stu, 1) //默认的条件是根据主键查询
	fmt.Println("主键查询：", stu)

	stu = Student{}
	DB.Take(&stu, "name = ?", "学生8") //根据指定条件查询
	fmt.Println("name = \"枫枫8\"：", stu)

	stuList := []Student{}
	target := "学生8' or 1=1;#"
	DB.Take(&stuList, fmt.Sprintf("name = '%s'", target)) //通过SQL注入获取数据库所有信息，因此最好不要使用字符串拼接方式进行SQL查询
	fmt.Println("SQL注入：", stuList)

}

func QueryByStruct(DB *gorm.DB) {
	var stu Student
	// stu.Name = "学生2"  //结构体查询方法，只有当主键有值时才会有用
	stu.ID = 2
	DB.Take(&stu) //如果结构体中主键字段已有数据，则会根据此值进行查询,因此查询前需要保证缓冲结构体的纯洁性
	fmt.Println("id = 2:", stu)
}

func QueryMulRecords(DB *gorm.DB) {
	var stuList []Student
	count := DB.Find(&stuList).RowsAffected //Find()用于多条查询,需要用切片保存结果数据，RowsAffected返回查询的条数
	fmt.Println("总共查询的结果数量为:", count)
	for _, stu := range stuList {
		fmt.Println(stu)
	}
	// 由于email是指针类型，所以看不到实际的内容
	// 但是json序列化之后，会转换为我们可以看得懂的方式
	data, _ := json.MarshalIndent(stuList, "", "")
	fmt.Println(string(data))
}

func ConditionalMulQuery(DB *gorm.DB) {
	var stuList []Student
	DB.Find(&stuList, []int{1, 4, 6}) //默认是根据主键查询
	fmt.Println(stuList)

	stuList = make([]Student, 0)
	DB.Find(&stuList, 1, 4, 6) //与上述的查询语句效果相同
	fmt.Println(stuList)

	stuList = make([]Student, 0)
	DB.Find(&stuList, "id in (?)", []int{1, 4, 6}) //条件查询
	fmt.Println(stuList)

	stuList = make([]Student, 0)
	DB.Find(&stuList, "name in (?)", []string{"学生1", "学生2"}) //条件查询
	fmt.Println(stuList)

}
