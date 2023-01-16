package four

import (
	"database/sql"
	"fmt"
)

func Transaction(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("创建事务失败，begin err:%v\n", err)
		return
	}

	_, err = tx.Exec("update `employee` set name = '琪琪' where id = ? ", 3)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql1 is failed,err:%v\n", err)
		return
	}
	_, err = tx.Exec("update `employee` set name = '赵云' where id = ?", 5)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql2 is failed,err:%v\n", err)
		return
	}
	tx.Commit()
	fmt.Println("事务执行成功！！")
}
