package main

import (
	"github.com/learn-go/test/loadblance/balance"
	"fmt"
	"math/rand"
)
func main() {
	insts := make([]*balance.Instance)
	for i := 0;i < 16;i++ {
		host := fmt.Sprintf("192.168.%d.%d",rand.Intn(255),rand.Intn(255))// sprintf不会打印结果，会以字符串返回

		one := balance.NewInstance(host, port)
		insts = append(insts)
	}
}
