package main

import "fmt"

func Adder() func(int) int {
	var x int
	return func(d int) int {
		// 这的d int 和 int 要和上面的对应上
			x += d
			return x
	}
}
func main()  {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
}
