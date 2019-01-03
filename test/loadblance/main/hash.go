package main

import (
	"github.com/learn-go/test/loadblance/balance"
	"hash/crc64"
	"fmt"
	"math/rand"
)
func init() {
	balance.RegisterBalancer("hash",&HashBalance{})
}
// 自己实现一个接口，使用hash算法
// 这样就支持三种算法了
type HashBalance struct {

}

func (p *HashBalance) DoBalance(insts []*balance.Instance,key ...string) (inst *balance.Instance,err error) {
	// key 可变参数，取key[0] 第0个参数
	var defkey string = fmt.Sprintf("%d",rand.Int())
	if (len(key)) > 0 {
		defkey = key[0]
	}

	lens := len(insts)
	if (lens ==0) {
		err = fmt.Errorf("No backend instance")
		return
	}
	crcTable := crc64.MakeTable(crc64.ECMA) // 固定用法
	hashval := crc64.Checksum([]byte(defkey),crcTable) // 把table传进去
	index := int(hashval)% lens  //类型uint64强制转成int
	inst = insts[index]
	return
}
