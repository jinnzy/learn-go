package main

import "fmt"

var pipe chan int

func add(a int,b int)  {
	var sum int
	sum = a + b

	pipe <- sum
	return
}

func main()  {

	pipe = make(chan int,1)
	//for i :=0;i < 100;i++{
	//	go test_print(i)
	//}
	//time.Sleep(10*time.Second)
	go add(2,5)
	sum :=<- pipe
	fmt.Println("sum=",sum)
}