package main

import (
	"fmt"
	"time"
)


func main()  {
	var i int

	go func() {
		i+=1
		fmt.Println(i)
	}()
	fmt.Println(i)
	time.Sleep(10 * time.Second)
}
