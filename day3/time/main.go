package main

import "fmt"

func modify(arr *[5]int)  {
	// arr *[5]int 声明一个指针类型的数组
	(*arr)[0] = 100 // 获取arr指向地址的值进行修改
	return
}
func main()  {
	var a [5]int
	modify(&a)
	for i := 0;i < len(a); i++{
		fmt.Println(a[i])
	}
}
