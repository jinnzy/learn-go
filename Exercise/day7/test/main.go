package main
import (
//	"fmt"
)
// 定义结构体
type PathError struct {
	Op   string
	Path string
	err string
}
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}
func test() error {
	return &PathError{
		Op:   "op",
		Path: "path",
	}
}
func main() {
	test()
}
