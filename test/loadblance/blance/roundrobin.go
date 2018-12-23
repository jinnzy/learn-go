package blance

import (
	"github.com/pkg/errors"
	"math/rand"
)
// 随机
type RoundRobinBalance struct {
	curIndex int // 存当前选择的主机
}
// 定义方法名为DoBalance
func (p *RoundRobinBalance) DoBalance(insts []*Instance) (inst *Instance,err error) {
	// 实例为0，报错返回err值
	lens := len(insts) // 传进来的主机数量
	if lens == 0 {
		err = errors.New("No interface")
		return
	}
	lens := len(insts)
	if (p.curl)

	return
}
