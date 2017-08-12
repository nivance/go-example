package main

import (
	"fmt"
	"os"
	//	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func init3() {
	//	age := "345"
	//	n := len(age)
	//	fmt.Println("result:" + string(age[0]))
	//	sa := make([]string, n)
	//	for i := 0; i < n; i++ {
	//		sa[i] = string(age[i])
	//		fmt.Println("age:" + string(age[i]))
	//	}
	//	fmt.Println("result:" + strings.Join(sa, ","))

	ReadAndWriteTag("C:\\Users\\lenovo\\Desktop\\7月第四周标签.xlsx", "C:\\Users\\lenovo\\Desktop\\a.out")
}

func ReadAndWriteTag(filePath string, resultFile string) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return
	}
	fd, _ := os.OpenFile(resultFile, os.O_RDWR|os.O_CREATE, 0644)
	defer fd.Close()
	sheet := xlFile.Sheet["Sheet1"]
	for _, row := range sheet.Rows {
		book_name := strings.Replace(row.Cells[0].Value, " ", "", -1)
		main_tag := strings.Replace(row.Cells[1].Value, " ", "", -1)
		sub_tags := strings.Replace(row.Cells[2].Value, " ", "", -1)
		age := strings.Replace(row.Cells[3].Value, " ", "", -1)

		n := len(age)
		sa := make([]string, n)
		for i := 0; i < n; i++ {
			sa[i] = string(age[i])
		}
		age = strings.Join(sa, ",")

		sql := "update t_book set TAGS = \"" + main_tag + "\", SUB_TAGS = \"" + sub_tags + "\", AGE = '" + age + "' where name = \"" + book_name + "\";\n"
		fd.WriteString(sql)
		fmt.Println(sql)
		sql = "update t_rec_action set DESCRIPTION = \"" + main_tag + "\" where name = \"" + book_name + "\";\n"
		fd.WriteString(sql)
		fmt.Println(sql)
	}

}
