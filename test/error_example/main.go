package main

import (
	"fmt"
	"io/ioutil"
	"github.com/pkg/errors"
)
// 定义错误的结构体

func main() {

	conent,err:= ioutil.ReadFile("test")
	if err !=nil{
		//错误处理
		fmt.Println(errors.Wrap(err,"open file failed"))
	}else {
		fmt.Println(string(conent))
	}
}