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

	mysql.GoInsert()
	logs.Logger.Info("-------Done.-------")

}
