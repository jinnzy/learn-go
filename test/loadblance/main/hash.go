package main

import (
	"github.com/learn-go/test/loadblance/balance"
		"fmt"
	"math/rand"
	"hash/crc32"
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
	crcTable := crc32.MakeTable(crc32.IEEE) // 固定用法
	hashval := crc32.Checksum([]byte(defkey),crcTable) // 把table传进去
	index := int(hashval)% lens  //crc64类型uint64强制转成int，可能会变成负数，这里换成crc32
	inst = insts[index]
	return
}
