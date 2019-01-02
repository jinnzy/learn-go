package main

import (
	"math/rand"
	"fmt"
)

func main()  {
	var n int
	n = rand.Intn(100) // 随机生成数字
	for {
		var input int
		fmt.Scanf("%d\n",&input) // 不加\n的时候命令行输入数字回车的话，回车也算一个字符。加上的话会匹配数字+\n 这时\n就会被忽略掉
		flag := false
		switch  {
		case input == n:
			fmt.Println("正确")
			flag = true
		case input > n:
			fmt.Println("大了")
		case input < n:
			fmt.Println("小了")
		}
		// 判断标志位，是true结束循环
		if flag {
			break
		}
	}
}