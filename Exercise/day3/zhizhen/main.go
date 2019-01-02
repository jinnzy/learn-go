package main

import "fmt"

func main()  {
	var test int = 10
	var p *int = &test // 将test的地址传给p，这样p就指向test的内存地址了
	fmt.Println(*p)  // *p  是取得p指向地址的值。
	*p = 100 // 修改指针指向地址的值
	fmt.Println(test) // test的值也会变成100
}
