package main

import "fmt"

//func swap(a *int,b *int)  {
//	// * 代表指针
//	tmp	:= *a // 取的指针所指向内存的值
//	*a = *b // 将指针的值进行交换
//	*b = tmp
//	return
//}
func swap(a int,b int) (int int) {

}

func main()  {
	first := 100
	second := 200
	first,second = second
	fmt.Println("first=",first)
	fmt.Println("second=",second)
}