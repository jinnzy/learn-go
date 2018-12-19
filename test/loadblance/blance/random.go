package blance

import (
	"github.com/pkg/errors"
	"math/rand"
)

type RandomBalance struct {

}
// 定义方法名为DoBalance
func (p *RandomBalance) DoBalance(insts []*Instance) (insts *Instance,err error) {
	// 实例为0，报错返回err值
	lens := len(insts)
	if lens == 0 {
		err = errors.New("No interface")
		return
		}
	index := rand.Intn(lens)
}
