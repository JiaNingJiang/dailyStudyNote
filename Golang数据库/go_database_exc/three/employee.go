package three

import (
	"database/sql"
	"fmt"
	"go_database_exc/two"
)

func PrepareQuery(db *sql.DB) {
	stmt, err := db.Prepare("select id,name,entrydate from `employee` where id >= ?")
	if err != nil {
		fmt.Printf("prepare failed , err:%v\n", err)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query(2)
	if err != nil {
		fmt.Printf("Query failed ,err:%v\n", err)
		return
	}
	defer rows.Close()

	u := new(two.User)
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Entrydate); err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id >= 2 :: id:%d,name:%s,entrydate:%s\n", u.ID, u.Name, u.Entrydate)
	}

	rows, err = stmt.Query(3)
	if err != nil {
		fmt.Printf("Query failed ,err:%v\n", err)
		return
	}
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Entrydate); err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id >= 3 :: id:%d,name:%s,entrydate:%s\n", u.ID, u.Name, u.Entrydate)
	}

}

func PrepareInsert(db *sql.DB) {
	stmt, err := db.Prepare("insert into `employee` (id,name) values (?,?)")
	if err != nil {
		fmt.Printf("prepare failed , err:%v\n", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(4, "李四")
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}

	_, err = stmt.Exec(5, "王五")
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}

	fmt.Println("插入成功...............")

}
