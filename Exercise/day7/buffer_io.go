package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)
type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}
func main() {
	file,err := os.Open("D:/123.txt")
	if err != nil {
		fmt.Println("read file err",err)
		return
	}
	defer file.Close() //关闭文件
	//var count CharCount
	reader := bufio.NewReader(file)
	for {
		// 循环读取文件内容，一直读到结尾，返回eof
		str,err := reader.ReadString('\n') // 读到行尾
		if err == io.EOF {
			// 读到文件尾部会返回EOF
			break // 跳出循环
		}else if err != nil {
			fmt.Println("read file err:",err)
		}
		b := strings.Trim(str,"\n")
		fmt.Println(b)

		//runeArr := []rune(str) // 生成字符数组
		//for _, v := range runeArr {
		//	switch {
		//	case v >= 'a' && v <= 'z': // a 对应ascii表示97，也可以写成>=97不过不直观
		//		fallthrough // 这个意思是匹配到这了不会结束还是往下走
		//	case v >= 'A' && v <= 'Z':
		//		count.ChCount++
		//	case v == ' ' || v== '\t':
		//		count.SpaceCount++
		//	case v >= '0' && v<= '9':
		//		count.NumCount++
		//	default:
		//		count.OtherCount++
		//	}
		//}
	}
	//fmt.Printf("char count:%d\n",count.ChCount)
	//fmt.Printf("num count:%d\n",count.NumCount)
	//fmt.Printf("space count:%d\n",count.SpaceCount)
	//fmt.Printf("other count:%d\n",count.OtherCount)
}