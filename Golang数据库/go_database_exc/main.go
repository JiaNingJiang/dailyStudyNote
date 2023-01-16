package main

import (
	"go_database_exc/five"
	"go_database_exc/one"
)

func main() {
	sql := one.InitDataBase()

	//u1 := two.QueryRow(sql)
	//
	//_ = u1

	//u2 := two.QueryMultiRow(sql)
	//_ = u2

	//two.InsertRow(sql)

	//two.UpdateRow(sql)

	//two.DeleteRow(sql)

	//three.PrepareQuery(sql)

	//three.PrepareInsert(sql)

	//four.Transaction(sql)

	five.UnKnowTable(sql)

	defer sql.Close()
}
