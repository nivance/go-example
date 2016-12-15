package main

import (
	"fmt"
	"strconv"

	"github.com/nivance/go-example/basic/logs"
	//	"github.com/nivance/go-example/basic/mysql"
)

func main() {
	//mysql.OrmInsert()
	//entity := mysql.OrmGetEntity(1)
	//entity := mysql.OrmGetAll()
	//logs.Logger.Info("-------Done.-------", entity)

	// mysql.UpdateTags("C:\\Users\\lenovo\\Desktop\\huiben\\11.7绘本标签.xlsx")
	//	logs.Logger.Info(mysql.ReadXlxs("c:\\tags.xlsx"))
	//mysql.QueryAllBooks()
	logs.Logger.Info("-------Done.-------")

	var i int = 7
	logs.Logger.Info(float64(i))
	a := fmt.Sprintf("%.4f", 3/float64(i))
	f, _ := strconv.ParseFloat(a, 2)
	logs.Logger.Info(f * 100)
	//	mysql.ReadNewTagExcel("d:\\newtags.xlsx")
	// mysql.ReadNewTagExcel("C:\\Users\\lenovo\\Desktop\\huiben\\youzan\\11.28.xlsx")
}
