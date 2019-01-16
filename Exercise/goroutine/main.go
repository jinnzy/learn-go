package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int , 10)
	ch2 := make(chan int,10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i*i

		}
	}()
	for {
		select {
		// select作用是防止阻塞，如 ch取完数据了，就会马上走default分支
		// 取不到数据就走默认分支，有数据的话在处理数据
		case v := <-ch:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <- time.After(time.Second):
			// 丢到值，time.After也是一个channel
			// 1秒没返回走这个分支，
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}
