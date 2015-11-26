package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nivance/go-example/basic/logs"
	"math/rand"
	"strconv"
	"time"
)

// use go-sql-driver

func GetGoDB() (db *sql.DB, err error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?autocommit=true")
}

func GoInsert() {
	db, err := GetGoDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 插入数据
	stmt, err := db.Prepare("insert into entries(title, content, created) values(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec("boss"+strconv.Itoa(rand.Intn(100)), "boss_command", time.Now())
	if err != nil {
		panic(err)
	}

	// 获取影响的行数
	affect, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	logs.Logger.Info("affected:", affect)

	// 获取自增id
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	logs.Logger.Info("id:", id)
}
