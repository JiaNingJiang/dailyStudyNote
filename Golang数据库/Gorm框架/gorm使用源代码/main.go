package main

import (
	"gorm.io/gorm"
	CustomDataType "gorm_exec/08.CustomDataType"
)

var DB *gorm.DB

func main() {
	//DB := _1_simpleConnect.InitGorm()
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // （日志输出的目标，前缀和日志包含的内容）
	//	logger.Config{
	//		SlowThreshold:             time.Second, // 慢 SQL 阈值
	//		LogLevel:                  logger.Info, // 日志级别
	//		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
	//		Colorful:                  true,        // 使用彩色打印
	//	},
	//)
	//
	//session := DB.Session(&gorm.Session{Logger: newLogger})
	//s1 := new(_1_simpleConnect.Student)
	//session.First(s1)
	//
	//s2 := new(_1_simpleConnect.Student)
	//DB.AutoMigrate(s2)

	//singleTable.CreateTable(DB)
	//singleTable.InsertSingle(DB)
	//singleTable.BatchInsert(DB)

	//singleTable.QuerySingle(DB)

	//singleTable.QuerySingleWithLog(DB)

	//singleTable.ConditionalQuery(DB)

	//singleTable.QueryByStruct(DB)

	//singleTable.QueryMulRecords(DB)

	//singleTable.ConditionalMulQuery(DB)

	//singleTable.SingleSaveUpdateAll(DB)

	//singleTable.SingleSaveUpdate(DB)

	//singleTable.BatchUpdateRow(DB)

	//singleTable.BatchUpdateRowsByStruct(DB)
	//singleTable.BatchUpdateRowsByMap(DB)

	//singleTable.DeleteWithQueryResult(DB)

	//singleTable.DeleteByCondition(DB)

	//hookUsing.ShowHookEffect(DB)

	//advancedQuery.CreateTable(DB)

	//advancedQuery.AdvancedQuery(DB)

	//advancedQuery.QueryWithStructOrMap(DB)

	//advancedQuery.QueryWithSelect(DB)

	//advancedQuery.OrderAndPageAndDistinct(DB)

	//advancedQuery.GroupQuery(DB)

	//advancedQuery.Subquery(DB)

	//advancedQuery.NamedParams(DB)

	//advancedQuery.CallChaining(DB)

	//MultiTableQuery.CreateOneToMul(DB)

	//MultiTableQuery.AddDataOneToMul(DB)

	//MultiTableQuery.AddDataOneToMulByForeignKey(DB)

	//MultiTableQuery.QueryOneToMul(DB)

	//MultiTableQuery.DeleteOneToMul(DB)

	//OneToOneTable.CreateOneToOneTable(DB)
	//
	//OneToOneTable.AddDateOneToOneTable(DB)

	//OneToOneTable.QueryOneToOneTable(DB)

	//OneToOneTable.DeleteOneToOneTable(DB)

	//MulToMulTabel.CreateTable(DB)

	//MulToMulTabel.AddData(DB)

	//MulToMulTabel.QueryData(DB)

	//MulToMulTabel.UpdateAndDelete(DB)

	//MulToMulTabel.CreateCustomTable(DB)

	//MulToMulTabel.InsertDataCustomTable(DB)

	//MulToMulTabel.QueryCustomTable(DB)

	//MulToMulTabel.CreateModelTable(DB)

	//MulToMulTabel.QueryArticlesByUser(DB)

	//CustomDataType.CreateTable(DB)

	//CustomDataType.InsertData(DB)

	//CustomDataType.QueryData(DB)

	//CustomDataType.CreateTableArr(DB)

	//CustomDataType.InsertDataArr(DB)

	//CustomDataType.QueryDataArr(DB)

	CustomDataType.MeiJu()
}
