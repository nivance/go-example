package mysql

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Entity struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}

func init() {
	orm.RegisterModel(new(Entity))
}
