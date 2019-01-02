package main

import "fmt"

func testSlice() {
	var slice []int
	var arr [5]int = [5]int{1,2,3,4,5}
	slice = arr[2:5] // 包含2,3的值 不包含4 要记住,包含起始的下标，不包含尾部的
	fmt.Println(slice)
	fmt.Println(len(slice)) // 长度是3
	fmt.Println(cap(slice)) // 容量是3
	slice = slice[0:1]
	fmt.Println(slice)
	fmt.Println(len(slice)) // 长度是1
	fmt.Println(cap(slice)) // 容量是3
}
func main()  {
	testSlice()
}
