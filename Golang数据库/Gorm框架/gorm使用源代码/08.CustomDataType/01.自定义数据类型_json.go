package CustomDataType

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Info struct {
	Status string `json:"status"`
	Addr   string `json:"addr"`
	Age    int    `json:"age"`
}

// 自定义的数据类型必须实现 Scanner 和 Valuer 接口，以便让 GORM 知道如何将该类型接收、保存到数据库，否则gorm不会将其存储到数据库中
// 如果是基本数据类型，就不需要实现 Scanner 和 Valuer 接口了

// Scan 从数据库中读取出来
func (i *Info) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	info := Info{}
	err := json.Unmarshal(bytes, &info)
	*i = info
	return err
}

// Value 存入数据库
func (i Info) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type User struct {
	ID   uint
	Name string
	Info Info `gorm:"type:string"`
}

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func InsertData(db *gorm.DB) {
	db.Create(&User{
		Name: "枫枫",
		Info: Info{
			Status: "牛逼",
			Addr:   "成都市",
			Age:    21,
		},
	})
}

func QueryData(db *gorm.DB) {
	var user User
	db.Take(&user)
	fmt.Println(user)
}
