package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int , 10)
	for i :=0;i<10;i++ {
		ch <- i
	}
	close(ch)  // 关闭ch管道
	for {
		var b int
		b,ok := <-ch // ok是为了检测channel关闭了，不检测的话塞入10个数据之后关闭，这边读取10个之后会一直是死循环
		// of - false说明管道已经关闭
		if ok == false {
			fmt.Println("chan is close")
			break
		}
		fmt.Println(b)
	}
}
