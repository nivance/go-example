package mysql

import (
	"database/sql"
	"strings"

	"bytes"
	"io"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nivance/go-example/basic/logs"
	"github.com/tealeg/xlsx"
	godrv "github.com/ziutek/mymysql/godrv"
)

const (
	driverName, tableName = "mysql", "t_book"
)

func ReadTags(limit int, offset int) (books []Book) {
	qb, _ := orm.NewQueryBuilder(driverName)
	qb.Select("NAME", "TAGS").From(tableName).OrderBy("ID").Limit(limit).Offset(offset)
	qbStr := qb.String()
	ormer := orm.NewOrm()
	ormer.Raw(qbStr).QueryRows(&books)
	logs.Logger.Info("results:", books)
	return books
}

func GetBookDB() (db *sql.DB, err error) {
	godrv.Register("SET NAMES UTF8")
	return sql.Open("mymysql", "tcp:192.168.199.224:3306*robot_test/root/Charles2015!")
}

func QueryBooks(limit int, offset int) (books []Book) {
	db, err := GetBookDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	qb, _ := orm.NewQueryBuilder(driverName)
	qb.Select("ID", "NAME", "TAGS").From(tableName).OrderBy("ID").Limit(limit).Offset(offset)
	logs.Logger.Info(qb.String())
	rows, err := db.Query(qb.String())
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	books = []Book{}
	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.Id, &book.Name, &book.Tags)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return books
}

func QueryAllBooks() {
	limit := 20
	offset := 0
	books := QueryBooks(limit, offset)
	leng := len(books)
	for leng == limit {
		for _, book := range books {
			logs.Logger.Info("result:", book.Id, book.Name, book.Tags)
		}
		offset = offset + 1
		books = QueryBooks(limit, offset*limit)
		leng = len(books)
	}
	for _, book := range books {
		logs.Logger.Info("result:", book.Id, book.Name, book.Tags)
	}
}

func UpdateTags(filePath string) {
	tags := ReadXlxs(filePath)
	limit := 20
	offset := 0
	db, err := GetBookDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	books := QueryBooks(limit, offset)
	leng := len(books)
	var buffer bytes.Buffer
	for leng == limit {
		for _, book := range books {
			//			logs.Logger.Info("result:", book.Id, book.Name, book.Tags)
			tag := tags[book.Name]
			if tag != "" {
				qb, _ := orm.NewQueryBuilder(driverName)
				qb.Update(tableName).Set("TAGS = '" + tag + "'").Where("NAME = '" + book.Name + "'")
				// db.Exec(qb.String())
				logs.Logger.Info(qb.String())
				buffer.WriteString(qb.String() + ";")
			}
		}
		offset = offset + 1
		books = QueryBooks(limit, offset*limit)
		leng = len(books)
	}
	write2File(buffer.String(), "d:\\updatetag.sql")
}

func write2File(data string, filename string) {
	if !checkFileIsExist(filename) {
		os.Create(filename)
	}
	fout, _ := os.OpenFile(filename, os.O_APPEND, 0777)
	io.WriteString(fout, data)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ReadXlxs(filePath string) (t map[string]string) {
	excelFileName := "c:\\tags.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return
	}
	var tags map[string]string = make(map[string]string)
	sheet := xlFile.Sheet["Sheet1"]
	for _, row := range sheet.Rows {
		name := strings.Replace(row.Cells[0].Value, " ", "", -1)
		tag := strings.Replace(row.Cells[1].Value, " ", "", -1)
		tag = strings.Replace(tag, "_x000D_", "", -1)
		if tag != "" && tag != "无" {
			tags[name] = tag
		}
	}
	return tags
}
