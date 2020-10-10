package main

import (
	"bufio"
	"fmt"
	"github.com/lexkong/log"
	"html/template"
	"io/ioutil"
	"os"
)

func test1() {
	file,err := os.Open("./tool_example/文件读取/filetoread.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	fileinfo,err := file.Stat()
	if err != nil {
		log.Fatal("123",err)
	}

	fileSize := fileinfo.Size()
	buffer := make([]byte,fileSize)

	bytesread,err := file.Read(buffer)
	if err != nil {
		log.Fatal("123",err)
	}
	fmt.Println("bytes read:",bytesread)
	fmt.Println("bytestream to string:",string(buffer))
}
func test2() {
	file,err := os.Open("./tool_example/文件读取/filetoread.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	rd := bufio.NewReader(file)
	buf := make([]byte,11000)
	_,err = rd.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

func test3()  {
	data,err := ioutil.ReadFile("./tool_example/文件读取/filetoread.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
func test4()  {
	t, err := template.ParseFiles().ParseFiles("e:/golang/go_pro/src/safly/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
func main()  {
	//test1()
	//test2()
	test3()
}
