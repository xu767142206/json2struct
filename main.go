package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var ctype string
var rpath string
var pname string
var structName string

func main() {
	//ePath, err := os.Executable()
	//if err != nil {
	//	log.Fatal(err)
	//}

	flag.StringVar(&ctype, "type", "json", "类型：json(暂只支持json)")
	flag.StringVar(&rpath, "path", "", "写入路径")
	flag.StringVar(&pname, "package", "main", "包名")
	flag.StringVar(&structName, "name", "demo", "结构名")
	flag.Parse()
	if rpath == "" {
		log.Fatal("path is no empty!")
	}

	//fmt.Println(toGoFieldCorrectName("ddddd"))

	byteJson, err := ioutil.ReadFile(rpath)
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := OpenFile(`./` + strings.Replace(structName, " ", "", -1) + `.go`)
	if err != nil {
		log.Fatal(err.Error())
	}
	stringWriter := bufio.NewWriter(file) //创建新的 Writer 对象
	//n41, err4 := w.Write([]byte("测试文件4字节流"))
	_, err = stringWriter.WriteString(`package ` + strings.Replace(pname, " ", "", -1) + "\n\n")
	if err != nil {
		log.Fatal(err.Error())
	}

	//stringWriter := &bytes.Buffer{}
	enc := NewEncoderWithNameAndTags(stringWriter,toGoFieldCorrectName(strings.Replace(structName, " ", "", -1)) , []string{"json"})

	if err := enc.Encode([]byte(byteJson)); err != nil {
		panic(err)
	}

	fmt.Println(getCurrentAbPathByCaller() + string(filepath.Separator) + strings.Replace(structName, " ", "", -1) + `.go`)

	stringWriter.Flush()
	file.Close()
}

// OpenFile 判断文件是否存在  存在则OpenFile 不存在则Create
func OpenFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return os.Create(filename) //创建文件
	}
	return os.OpenFile(filename, os.O_APPEND, 0777) //打开文件
}


// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}