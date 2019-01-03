package main

import (
	"github.com/learn-go/test/loadblance/balance"
	"fmt"
	"math/rand"
	"time"
)
func main() {
	insts := make([]*balance.Instance,0) // 声明一个空切片
	for i := 0;i < 16;i++ {
		host := fmt.Sprintf("192.168.%d.%d",rand.Intn(255),rand.Intn(255))// sprintf不会打印结果，会以字符串返回
		one := balance.NewInstance(host, 8080)
		insts = append(insts,one)
	}
	var balanceName = "hash"
	//fmt.Println(insts[0])
	//balancer := &balance.RoundRobinBalance{} // 轮询
	//balancer := &balance.RandomBalance{} // 随机
	for {
		inst,err := balance.DoBalance(balanceName,insts)
		if err !=nil{
			fmt.Println("do balance err:",err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
