package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/nivance/go-example/basic/logs"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
)

type Entity struct {
	Id      int `PK`
	Title   string
	Content string
	Created time.Time
}

func init() {
	orm.RegisterDriver("mymysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mymysql", "tcp:127.0.0.1:3306*test/root/")
	orm.Debug = true
	orm.RegisterModel(new(Entity))
	orm.BootStrap()
}

func GetAll() (entitys []Entity) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id", "title", "content", "created").From("entity")
	sql := qb.String()
	ormer := orm.NewOrm()
	ormer.Raw(sql).QueryRows(&entitys)
	logs.Logger.Info("results:", entitys)
	return entitys
}

func GetBlog(id int) (blog Entity) {
	blog.Id = id
	ormer := orm.NewOrm()
	ormer.Read(&blog)
	logs.Logger.Info("results:", blog)
	return blog
}

func SaveBlog(blog Entity) {
	o := orm.NewOrm()
	o.Using("default")
	blog.Created = time.Now()
	o.Insert(&blog)
}

func DelBlog(blog Entity) {
	o := orm.NewOrm()
	affectnum, err := o.QueryTable("entity").Filter("id", blog.Id).Delete()
	if err != nil {
		logs.Logger.Error("delete failed", err)
	}
	logs.Logger.Info("delete success, affectnum:", affectnum)
	return
}
