package mysql

import (
	"database/sql"
	"github.com/nivance/go-example/basic/logs"
	godrv "github.com/ziutek/mymysql/godrv"
	"math/rand"
	"strconv"
	"time"
)

// use mymysql driver
func GetDB() (db *sql.DB, err error) {
	//	models := []mysql.Entity{}
	//	model := mysql.Entity{1, "2", "3", time.Now()}
	//	models = append(models, model)
	godrv.Register("SET NAMES UTF8")
	return sql.Open("mymysql", "tcp:127.0.0.1:3306*test/root/")
}

func Insert() {
	db, err := GetDB()
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

func Query() (models []Entity) {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//rows, err = db.Query("select * from entries where id = " + strconv.Itoa(id))
	rows, err := db.Query("select * from entries")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	models = []Entity{}
	for rows.Next() {
		model := Entity{}
		err = rows.Scan(&model.Id, &model.Title, &model.Content, &model.Created)
		if err != nil {
			panic(err)
		}
		logs.Logger.Info("model:", model)
		models = append(models, model)
	}
	logs.Logger.Info("models:", models)
	return models
}
