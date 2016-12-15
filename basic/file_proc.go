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

const destDir = "C:\\Users\\lenovo\\Desktop\\huiben\\cover"
const jsonDir = "C:\\Users\\lenovo\\Desktop\\huiben\\json_file"

func init() {
	var srcDir = "C:\\Users\\lenovo\\Desktop\\huiben\\1214"
	// var srcDir = "C:\\Users\\lenovo\\Desktop\\huiben\\cover\\audio"
	checkDirExist(destDir, true)
	if !checkDirExist(srcDir, false) {
		fmt.Println("有文件夹不存在")
		return
	}
	// filepath.Walk(srcDir, walkDir)
	filepath.Walk(srcDir, walkFunc)
	// supplyFile(jsonDir, srcDir)
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
		//		if strings.Contains(info.Name(), "json") {
		//			fmt.Println("处理文件:::" + path)
		//			data, err := ioutil.ReadFile(path)
		//			if err != nil {
		//				panic(err)
		//			}
		//			data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
		//			data, _ = GbkToUtf8(data)
		//			var book Book
		//			if err := json.Unmarshal(data, &book); err != nil {
		//				panic(err)
		//			}
		//			var bookName = book.Title
		//			var fileName = strings.Split(info.Name(), ".")[0]
		//			var zipFileName = strings.Split(book.ImageZipfile, ".")[0]
		//			if !strings.EqualFold(fileName, bookName) || !strings.EqualFold(zipFileName, bookName) {
		//				fmt.Println(info.Name() + ":::文件名:[" + bookName + "], zip文件名:[" + zipFileName + "]有误")
		//				fmt.Println(fileName + ":::文件名")
		//				fmt.Println(bookName + ":::书名")
		//				fmt.Println(zipFileName + ":::zip文件名")
		//				panic(info.Name())
		//			}
		//			fmt.Printf("%s 格式正确\n", path)
		//		}
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
	if err != nil {
		panic(err)
	}
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
