package main

import "fmt"

func test(n int) int {
	if n == 1 {
		return 1
	}
	return test(n -1) * n
}
func main()  {
	m := test(5)
	fmt.Println(m)
}
