// basic project main.go
package main

import (
	"github.com/nivance/go-example/basic/logs"
	"github.com/nivance/go-example/basic/mysql"
)

func main() {
	// use mymysql driver
	//	mysql.Insert()
	//	entities := mysql.Query()
	//	logs.Logger.Debug("result:", entities)

	//mysql.OrmInsert()
	// entity := mysql.OrmGetEntity(1)
	entity := mysql.OrmGetAll()
	logs.Logger.Info("-------Done.-------", entity)

}
