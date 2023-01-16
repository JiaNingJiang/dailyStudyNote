package one

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitDataBase() *sql.DB {
	db, err := sql.Open("mysql", "root:2486@tcp(127.0.0.1:3306)/itcast")
	if err != nil {
		fmt.Println(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("数据库连接失败,err:", err)
	}

	return db
}
