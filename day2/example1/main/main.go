package main

import "fmt"

func testSlice() {
	var a [5]int = [...]int{1,2,3,4,5}
	s := a[1:]
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 11)
	fmt.Println(s)
}

func main()  {
	testSlice()
}