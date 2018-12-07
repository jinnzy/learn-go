package main

import "fmt"

func main()  {
	str := "hello, world,中国"

	for index,val := range str{
		fmt.Printf("index[%d] val[%c] len[%d]\n",index,val,len([]byte(string(val))))
		// len([]byte(string(val))) 先把val转成string类型，再转成byte进行统计
	}
}
