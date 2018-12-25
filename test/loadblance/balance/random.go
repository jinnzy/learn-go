package balance

import (
	"github.com/pkg/errors"
	"math/rand"
)
// 随机
type RandomBalance struct {

}
// 定义方法名为DoBalance
func (p *RandomBalance) DoBalance(insts []*Instance) (inst *Instance,err error) {
	// 实例为0，报错返回err值
	lens := len(insts) // 传进来的主机数量
	if lens == 0 {
		err = errors.New("No interface")
		return
		}
	index := rand.Intn(lens) // 生成随机数
	inst = insts[index]  // 找到对应主机返回，这样就是随机值了
	return
	}
