package models

import (
	"database/sql"
	"github.com/astaxie/beedb"
	log "github.com/cihub/seelog"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
)

type Blog struct {
	Id      int `PK`
	Title   string
	Content string
	Created time.Time
}

func GetLink() beedb.Model {
	db, err := sql.Open("mymysql", "blog/astaxie/123456")
	log.Info("GetLink:", db)
	if err != nil {
		log.Critical("error", err)
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}

func GetAll() (blogs []Blog) {
	db := GetLink()
	log.Info(db)
	db.FindAll(&blogs)
	return
}

func GetBlog(id int) (blog Blog) {
	db := GetLink()
	db.Where("id=?", id).Find(&blog)
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	db := GetLink()
	db.Save(&blog)
	return bg
}

func DelBlog(blog Blog) {
	db := GetLink()
	db.Delete(&blog)
	return
}
