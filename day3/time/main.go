package main

import "fmt"

func test(a int,b int) int{
	result := func(a1 int,b1 int) int {
		// 定义一个匿名函数
		return a1 + b1
	}(a,b) // 定义的时候直接调用，定义完这样调用result(a,b)

	return result
}
func main()  {
	fmt.Println(test(100,200))
}
