package mysql

import (
	"time"
)

type Entity struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}
