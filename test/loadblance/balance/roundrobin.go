package balance

import (
	"github.com/pkg/errors"
)
// 轮询
func init() {
	// 程序启动就会把随机算法注册到mgr里
	RegisterBalancer("roundRobin",&RoundRobinBalance{})
}
type RoundRobinBalance struct {
	curIndex int // 存当前选择的主机
}
// 定义方法名为DoBalance
func (p *RoundRobinBalance) DoBalance(insts []*Instance,key ...string) (inst *Instance,err error) {
	// 实例为0，报错返回err值
	lens := len(insts) // 传进来的主机数量
	if lens == 0 {
		err = errors.New("No interface")
		return
	}
	//fmt.Println(p.curIndex)
	if p.curIndex >= lens { // 大于等于是因为数组最大下标是长度-1，所以找不到对应值把curlIndex设置为0
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens

	//p.curIndex = (p.curIndex + 1) % lens # + 1 取余，和上面的类似都是防止越界
	return
}
