package main

import "fmt"

func print(n int)  {
	for i := 1;i < n+1 ;i++ {
		for j :=0;j < i;j++{
			fmt.Printf("A") // 格式化输出
		}
		fmt.Println() // 代表换行，
	}
}
func main()  {
	print(4)
}