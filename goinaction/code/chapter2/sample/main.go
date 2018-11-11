package sample

import (
	"log"
	"os"

	_ "github.com/learn-go/goinaction/code/chapter2/sample/matchers"
	"github.com/learn-go/goinaction/code/chapter2/sample/search"
)


// init在main之前调用
func init()  {
	log.SetOutput(os.Stdout)  // 将日志输出到标准输出
}

// main是整个程序的入口
func main()  {
	// 使用特定的项做搜索
	search.Run("president")
}