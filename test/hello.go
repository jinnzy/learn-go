package main

import "fmt"


// c chan int 第三个变量是传入管道
func add(a int,b int,c chan int)  {
	var sum int
	sum = a + b

	c <- sum // 2. 将结果传入管道中
	return
}

func main()  {
	// 入口函数
	var pipe chan int 	// 声明管道类型
	pipe = make(chan int,1)  // 声明管道大小
	go add(2,5,pipe) // 1. goroute并发调用add函数
	sum :=<- pipe // 3. 从管道中取出结果
	fmt.Println("sum=",sum)
}