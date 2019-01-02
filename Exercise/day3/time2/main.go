package main

import (
	"time"
	"fmt"
)

func test()  {
	time.Sleep(time.Second * 1)
}

func main()  {
	now := time.Now()
	fmt.Println(now)

	start := time.Now().UnixNano() // 纳秒 最少的单位
	test()
	end := time.Now().UnixNano()
	fmt.Printf("cost:%d",(end - start)/1000000000) // 转换成秒
}