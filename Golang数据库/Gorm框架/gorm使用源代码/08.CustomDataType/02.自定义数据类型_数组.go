package CustomDataType

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Ports []string

// Scan 从数据库中读取出来
func (p *Ports) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	*p = strings.Split(string(bytes), "|") //以 | 为分隔符，将字符串切成字符串数组
	return nil
}

// Value 存入数据库
func (p Ports) Value() (driver.Value, error) {
	return strings.Join(p, "|"), nil //给Ports数组各个元素之间添加一个 | 分隔符
}

type Address struct {
	ID    uint
	IP    string
	Ports Ports `gorm:"type:string"`
}

func CreateTableArr(db *gorm.DB) {
	db.AutoMigrate(&Address{})
}
func InsertDataArr(db *gorm.DB) {
	db.Debug().Create(&Address{
		IP: "192.168.2.2",
		Ports: Ports{
			"80",
			"8080",
		},
	})
}

func QueryDataArr(db *gorm.DB) {
	var addr Address
	db.Take(&addr)
	fmt.Println(addr)
}
