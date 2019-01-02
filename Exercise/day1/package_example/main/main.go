package main

import (
	"github.com/learn-go/Exercise/day1cise/day1/package_example/calc"
	"fmt"
)

func main()  {
	sum := calc.Add(100,300)
	sub := calc.Sub(100,300)

	fmt.Println(sum)
	fmt.Println(sub)

}
