package main

import (
	"github.com/learn-go/day1/goroute_example/goroute"
	"fmt"
)

func main()  {
	//var pipe chan int ，用 := 就不用定义pipe变量了
	pipe := make(chan int,2)
	go goroute.Add(2,3,pipe)
	sum :=  <- pipe

	fmt.Println(sum)
}
