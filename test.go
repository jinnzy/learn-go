package main

import (
	"fmt"
)

func test_goroute(a int, b int) {
	sum := a + b // 自定初始化变量类型，根据a 和 b的类型自动生成
	fmt.Println(sum)
}
func test_print(a int) {
	fmt.Println(a)
}
