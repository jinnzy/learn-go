package main

import "fmt"

func add(a int,arg...int) int {
	//var c int
	for _,val := range arg{
		a = val + a
	}
	return a
}
func main()  {
	a := add(1,2,4,4,5,6,6*)
	fmt.Println(a)
}
