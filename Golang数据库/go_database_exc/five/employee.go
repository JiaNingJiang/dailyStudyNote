package five

import (
	"database/sql"
	"fmt"
)

func UnKnowTable(db *sql.DB) {

	stmt, err := db.Prepare("select id,name,entrydate from `employee` where id = ?")
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

	cols, err := rows.Columns()

	for index, col := range cols {
		fmt.Printf("employee表的第%d列名为:%s\n", index, col)
	}

	vals := make([]interface{}, len(cols))
	for i, _ := range cols {
		vals[i] = new(sql.RawBytes)
	}

	for rows.Next() {
		if err = rows.Scan(vals...); err != nil {
			fmt.Println("row scan is err:", err)
		}
		for i, _ := range vals {
			fmt.Printf("Query获取的数据————%s:%s\n", cols[i], vals[i])
		}
	}
}
