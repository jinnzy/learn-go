package main

import (
		"time"
	"fmt"
)

const (
	Man = 1
	Female = 2

)

func main()  {
	second := time.Now().Unix()
	if (second % Female == 0 ){
		fmt.Println("female")
	}else {
		fmt.Println("man")
	}
}