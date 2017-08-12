package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	//	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const destDir = "C:\\Users\\lenovo\\Desktop\\huiben\\outer"
const jsonDir = "C:\\Users\\lenovo\\Desktop\\huiben\\json_file"

func init() {
	//	var srcDir = "C:\\Users\\lenovo\\Desktop\\huiben\\ready\\8月第二周\\本周新增zip"
	// var srcDir = "C:\\Users\\lenovo\\Desktop\\huiben\\ready\\替换英语音频的"
	var srcDir = "C:\\Users\\lenovo\\Desktop\\huiben\\ready\\etcb"
	checkDirExist(destDir, true)
	if !checkDirExist(srcDir, false) {
		fmt.Println("有文件夹不存在")
		return
	}
	//	filepath.Walk(srcDir, walkDir)
	//	filepath.Walk(srcDir, walkFunc)
	//	supplyFile(jsonDir, srcDir)

	// 制作英文版音频etcb文件
	//	filepath.Walk(srcDir, produceEnAudio)
	//	buf.WriteString("update t_audio a, t_book b set a.IMAGE = b.IMAGE where a.BOOK_ID = b.ID and length(a.IMAGE) = 0;")
	//	fmt.Println(buf.String())

	// 重命名音频文件
	//	filepath.Walk(srcDir, walkAndRenameFile)

	// 生成更新etcb sql
	// filepath.Walk(srcDir, buildEtcbSQL)
}

type Book struct {
	Age          string
	Auther       string
	BriefIntro   string
	ContentType  int
	ImageZipfile string
	Isbn         string
	Language     int
	PermitPage   int
	Press        string
	Price        int
	Reader       string
	Subject      string
	Tags         string
	Time         string
	Title        string
	TotalPage    int
	Version      string
}

func walkDir(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		if strings.ContainsAny(info.Name(), " ") {
			panic(info.Name() + "  有空格。。。。。。。。。。。。")
		}
		fmt.Printf("%s\n", info.Name())
	}
	//	else {
	//		fmt.Printf("%s\n", info.Name())
	//	}
	return nil
}

/*
*copy 缺少json文件的绘本到目标目录
 */
func supplyFile(jsonPath string, forDir string) {
	var fileMap map[string]string = make(map[string]string)
	_ = filepath.Walk(jsonPath, func(filename string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if _, ok := fileMap[info.Name()]; !ok {
				var value = jsonPath + "\\" + info.Name()
				fileMap[info.Name()] = value
			}
		}
		return nil
	})
	_ = filepath.Walk(forDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			var fileName = path + "\\" + info.Name() + ".json"
			_, err := os.Stat(fileName)
			if err == nil || os.IsExist(err) {
				fmt.Println(fileName + " is exist")
			} else {
				if v, ok := fileMap[info.Name()+".json"]; ok {
					// fmt.Println("############# " + v)
					if _, err := CopyFile(fileName, v); err != nil {
						panic(err)
					}
				} else {
					fmt.Println("~~~~~~~~~~~~~ " + fileName + " is not exist")
				}
			}
		}
		return nil
	})

	//fmt.Printf("%s\n", fileMap)
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		if strings.Contains(info.Name(), "json") {
			fmt.Println("处理文件:::" + path)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
			data, _ = GbkToUtf8(data)
			var book Book
			if err := json.Unmarshal(data, &book); err != nil {
				panic(err)
			}
			var bookName = book.Title
			var fileName = strings.Split(info.Name(), ".")[0]
			var zipFileName = strings.Split(book.ImageZipfile, ".")[0]
			if !strings.EqualFold(fileName, bookName) || !strings.EqualFold(zipFileName, bookName) {
				fmt.Println(info.Name() + ":::文件名:[" + bookName + "], zip文件名:[" + zipFileName + "]有误")
				fmt.Println(fileName + ":::文件名")
				fmt.Println(bookName + ":::书名")
				fmt.Println(zipFileName + ":::zip文件名")
				panic(info.Name())
			}
			fmt.Printf("%s 格式正确\n", path)
		}
		if _, err := CopyFile(destDir+string(os.PathSeparator)+info.Name(), path); err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("%s\n", path)
	}
	return nil
}

func walkAndFindJsonFile(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		if strings.Contains(info.Name(), "json") {
			fmt.Println("处理文件:::" + path)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
			data, _ = GbkToUtf8(data)
			var book Book
			if err := json.Unmarshal(data, &book); err != nil {
				panic(err)
			}
			var bookName = book.Title
			var fileName = strings.Split(info.Name(), ".")[0]
			var zipFileName = strings.Split(book.ImageZipfile, ".")[0]
			if !strings.EqualFold(fileName, bookName) || !strings.EqualFold(zipFileName, bookName) {
				fmt.Println(info.Name() + ":::文件名:[" + bookName + "], zip文件名:[" + zipFileName + "]有误")
				fmt.Println(fileName + ":::文件名")
				fmt.Println(bookName + ":::书名")
				fmt.Println(zipFileName + ":::zip文件名")
				panic(info.Name())
			}
			fmt.Printf("%s 格式正确\n", path)
		}
		if _, err := CopyFile(destDir+string(os.PathSeparator)+info.Name(), path); err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("%s\n", path)
	}
	return nil
}

var bookId = ""
var extCode = "_en" // 制作英文etcb文件时用
// var extCode = "0" 	// 音频替换时用

func walkAndRenameFile(path string, info os.FileInfo, err error) error {
	fmt.Printf("%s\n", path)
	// fmt.Printf("bookId:%s\n", bookId)
	if !info.IsDir() {
		var buf bytes.Buffer
		var names = strings.Split(info.Name(), "_")
		for i, v := range names {
			if len(buf.String()) > 0 {
				buf.WriteString("_")
			}
			if i == 0 {
				buf.WriteString(bookId)
				buf.WriteString(v)
				buf.WriteString(extCode)
			} else {
				buf.WriteString(v)
			}
		}
		os.Rename(path, strings.Replace(path, info.Name(), buf.String(), -1))
		fmt.Printf("%s\n", strings.Replace(path, info.Name(), buf.String(), -1))
		fmt.Printf("%s\n", buf.String())

	} else {
		fmt.Printf("%s\n", path)
	}
	return nil
}

/*
* etcb update sql
 */
func buildEtcbSQL(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		var names = strings.Split(info.Name(), ".")
		var ns = strings.Split(names[0], "_")
		fmt.Printf("update t_book set link = \"http://oarfc773f.bkt.clouddn.com/" + info.Name() + "\", version = version + 1  where name = \"" + ns[1] + "\";")
		fmt.Println()
		fmt.Printf("update t_audio set url = \"http://oarfc773f.bkt.clouddn.com/" + info.Name() + "\" where name = \"" + ns[1] + "\" and AUTHENTICATE = 1 and status = 1;")
		fmt.Println()
	} else {
		fmt.Printf("%s\n", path)
	}
	return nil
}

var etcbContent = ""
var buf bytes.Buffer // 英文etcb insert sql

/*
* 制作英语版音频
 */
func produceEnAudio(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && strings.Contains(path, "etcb") { // 只读etcb文件
		// fmt.Printf("%s\n", path)
		// fmt.Printf("%s\n", info.Name())
		var names = strings.Split(info.Name(), ".") // 文件名：100000104ESpPOMO_好安静的蟋蟀.etcb
		var ns = strings.Split(names[0], "_")
		var bookName = ns[1]
		bookId = ns[0][0:9]
		var audioDir = strings.Replace(path, info.Name(), "", -1) + bookName
		filepath.Walk(audioDir, walkAndRenameFile) //重命名音频文件
		// 替换etcb音频文件url
		data, err := ioutil.ReadFile(path) // 读取Etcb内容
		if err != nil {
			panic(err)
		}
		etcbContent = string(data)
		// fmt.Println(etcbContent)
		filepath.Walk(audioDir, replaceAudioString)
		// fmt.Println("result:" + etcbContent)
		var enFileName = names[0] + extCode + "." + names[1]
		var resultFileName = strings.Replace(path, info.Name(), "", -1) + enFileName
		write2File(resultFileName, etcbContent)
		buf.WriteString("insert into t_audio(BOOK_ID, NAME, URL, AUTHOR, AUTHENTICATE, TYPE, STATUS, PLAY_TIMES, CREATE_TIME) values(" + bookId + ", \"" + bookName + "\", \"http://oarfc773f.bkt.clouddn.com/" + enFileName + "\", \"嘟巴英文版\", 1, 0, 1, FLOOR(5 + (RAND() * 8)), now());\n")
	}
	return nil
}

/*
* 替换音频url:
*	读取文件名，截取相同部分(如1_1_r)，替换
 */
func replaceAudioString(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		var subName = info.Name()[strings.IndexAny(info.Name(), "_")+4 : len(info.Name())] // 文件名：100000104tvqc_en_1_1_r.mp3，取2_10_a.mp3
		// fmt.Println("subName:" + subName)
		// fmt.Println("etcbContent:" + etcbContent)
		// 取etcb文件对应MP3字符串
		var lastSubNameIndex = strings.Index(etcbContent, subName)
		if lastSubNameIndex > 0 {
			// fmt.Println("etcbContent-sub1:" + etcbContent[0:lastSubNameIndex])
			var oldMp3 = etcbContent[strings.LastIndex(etcbContent[0:lastSubNameIndex], "/")+1 : lastSubNameIndex+len(subName)]
			// fmt.Println("oldMp3:" + oldMp3)
			etcbContent = strings.Replace(etcbContent, oldMp3, info.Name(), -1)
		} else {
			fmt.Println("额外的音频：" + info.Name())
		}
	}
	return nil
}

func write2File(path, content string) {
	file, err := os.Create(path)
	CheckError(err)
	defer file.Close()
	_, err1 := file.WriteString(content)
	CheckError(err1)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func walkFuncPrintFile(path string, info os.FileInfo, err error) error {
	fmt.Println("http://oarfc773f.bkt.clouddn.com/" + info.Name())
	return nil
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	fmt.Println("copy to:::" + dstName)
	src, err := os.Open(srcName)
	CheckError(err)
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func checkDirExist(dirPath string, create bool) bool {
	var exist = true
	if _, err := os.Open(dirPath); os.IsNotExist(err) {
		if create {
			if err := os.MkdirAll(dirPath, 0777); err != nil {
				fmt.Println(err)
				exist = false
			}
		}
	}
	return exist
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
