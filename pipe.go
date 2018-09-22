package main

import (
	"fmt"
)

func test_pipe() {
	pipe := make(chan int, 3) // chan是管道，创建管道容量为3，int类型
	pipe <- 1                 // 把 1 放到管道里去
	pipe <- 2                 // 把 2 放到管道里去
	fmt.Println(len(pipe))
}
