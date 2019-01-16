package timer

import (
	"time"
	"fmt"
)

func main()  {
	for {
		t := time.NewTicker(time.Second)

		select {
		case <- t.C:
			fmt.Println("超时")
		}
		t.Stop() // 用完time要关掉，否则会发生内存泄漏
	}
}