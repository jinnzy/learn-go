package balance

import (
	"github.com/pkg/errors"
	"math/rand"
)
// 随机
func init() {
	// 程序启动就会把随机算法注册到mgr里
	RegisterBalancer("random",&RandomBalance{})
}
type RandomBalance struct {

}
// 定义方法名为DoBalance
func (p *RandomBalance) DoBalance(insts []*Instance,key ...string) (inst *Instance,err error) {
	// 实例为0，报错返回err值，key ...string为合并参数可用可不用
	lens := len(insts) // 传进来的主机数量
	if lens == 0 {
		err = errors.New("No interface")
		return
		}
	index := rand.Intn(lens) // 生成随机数
	inst = insts[index]  // 找到对应主机返回，这样就是随机值了
	return
	}
