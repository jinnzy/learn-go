package main

import (
	"time"
)

func test() {
	defer func() {
		// 捕获panic
		if err := recover();err != nil {
			fmt.Println("panic err:",err)
		}
	}()

	var m map[string]int
	m["stu"] = 100 // 会panic，访问空的map
}
func main() {
	for i := 0;i<10;i++ {
		go test()
	}
	time.Sleep(time.Second*1000)
}