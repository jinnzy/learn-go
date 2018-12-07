package main

import "fmt"

func main()  {
	var a int = 20
	switch a {
	case 0:
		fmt.Println("a is equal 0")
		fmt.Println("yes")
		fallthrough // 这个意思是匹配到这了不会结束还是往下走
	case 10,20,30: // 匹配到10 20 30 会执行这个分支
		fmt.Println("a is equal 10 20 30")
	default:
		fmt.Println("a is equal default")
	}
}