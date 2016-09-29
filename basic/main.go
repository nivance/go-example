package main

import (
	"github.com/nivance/go-example/basic/logs"
	"github.com/nivance/go-example/basic/mysql"
)

func main() {
	//mysql.OrmInsert()
	//entity := mysql.OrmGetEntity(1)
	//entity := mysql.OrmGetAll()
	//logs.Logger.Info("-------Done.-------", entity)

	mysql.UpdateTags("c:\\tags.xlsx")
	//	logs.Logger.Info(mysql.ReadXlxs("c:\\tags.xlsx"))
	//mysql.QueryAllBooks()
	logs.Logger.Info("-------Done.-------")
}
