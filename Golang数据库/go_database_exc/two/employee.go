package two

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID        int
	WorkNo    string
	Name      string
	Gender    string
	Age       int
	IdCard    string
	Entrydate string
}

func QueryRow(db *sql.DB) *User {

	u := new(User)
	err := db.QueryRow("select id,name,entrydate from `employee` where id = ?", 1).Scan(&u.ID, &u.Name, &u.Entrydate)
	if err != nil {
		fmt.Println("数据库读取失败,err:", err)
		return nil
	}
	fmt.Printf("id:%d,name:%s,entrydate:%s\n", u.ID, u.Name, u.Entrydate)
	return u
}

func QueryMultiRow(db *sql.DB) []User {
	rows, err := db.Query("select id,name,age from `employee` where age between ? and ?", 18, 30)
	if err != nil {
		fmt.Println("数据库读取失败,err:", err)
		return nil
	}
	defer rows.Close()

	u := new(User)
	users := make([]User, 0)
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			fmt.Printf("row scan is failed,err:%v\n", err)
		}
		fmt.Printf("id:%d,name:%s,age:%d\n", u.ID, u.Name, u.Age)
		users = append(users, *u)
	}

	return users
}

func InsertRow(db *sql.DB) {
	ret, err := db.Exec("insert into `employee` (id,workno,name) values (?,?,?)", 4, 4, "tick")
	if err != nil {
		fmt.Println("数据插入失败,err:", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("Get LastInsert ID failed,err:", err)
		return
	}
	fmt.Printf("insert success,the id is %d\n", id)
}

func UpdateRow(db *sql.DB) {
	ret, err := db.Exec("update `employee` set name = ? where id = ?", "李红", 1)
	if err != nil {
		fmt.Println("数据更新失败,err:", err)
		return
	}
	n, err := ret.RowsAffected() //此次Exec操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected is failed,err:%v\n", err)
		return
	}
	fmt.Printf("update success,Affected row:%d\n", n)
}

func DeleteRow(db *sql.DB) {
	ret, err := db.Exec("delete from `employee` where id = ?", 4)
	if err != nil {
		fmt.Println("数据删除失败,err:", err)
		return
	}
	n, err := ret.RowsAffected() //此次Exec操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected is failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete success,Affected row:%d\n", n)
}
