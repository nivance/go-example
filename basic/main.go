package main

import (
	"fmt"
	//	"io/ioutil"
	"strconv"
	//	"strings"

	"github.com/nivance/go-example/basic/logs"
	//	"github.com/nivance/go-example/basic/mysql"
)

func main() {
	//mysql.OrmInsert()
	//entity := mysql.OrmGetEntity(1)
	//entity := mysql.OrmGetAll()
	//logs.Logger.Info("-------Done.-------", entity)

	// mysql.UpdateTags("C:\\Users\\lenovo\\Desktop\\huiben\\11.7绘本标签.xlsx")
	//	logs.Logger.Info(mysql.ReadXlxs("c:\\tags.xlsx"))
	//mysql.QueryAllBooks()
	logs.Logger.Info("-------Done.-------")

	var i int = 7
	logs.Logger.Info(float64(i))
	a := fmt.Sprintf("%.4f", 3/float64(i))
	f, _ := strconv.ParseFloat(a, 2)
	logs.Logger.Info(f * 100)
	//	mysql.ReadTagsFromFile("C:\\Users\\lenovo\\Desktop\\20170422.xlsx")
	//	mysql.ReadNewTagExcel("d:\\newtags.xlsx")
	//mysql.ReadNewTagExcel("C:\\Users\\lenovo\\Desktop\\huiben\\youzan\\0105.xlsx")

	/*data, _ := ioutil.ReadFile("C:\\Users\\lenovo\\Desktop\\huiben\\ready\\替换英语音频的\\100000104ESpPOMO_好安静的蟋蟀.etcb") // 读取Etcb内容
	var source = string(data)
	var pattern string = "3_2_a.mp3"
	fmt.Println("index : %%d", strings.Index(source, pattern))
	fmt.Println(source[0:strings.Index(source, pattern)])
	fmt.Println(source[strings.LastIndex(source[0:strings.Index(source, pattern)], "/")+1 : strings.Index(source, pattern)+len(pattern)])
	fmt.Println(source[strings.LastIndex(source[0:strings.Index(source, pattern)], "/")+1 : strings.Index(source, pattern)+len(pattern)])*/
}
