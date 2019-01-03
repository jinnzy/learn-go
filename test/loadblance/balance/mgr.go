package balance

import "fmt"

type BalanceMgr struct {
	allBalancer map[string]Balancer
}
var mgr = BalanceMgr{
	allBalancer:make(map[string]Balancer),
}

func (p *BalanceMgr) registerBalancer(name string,b Balancer) {
	p.allBalancer[name] = b
}

func RegisterBalancer(name string,b Balancer) {
	// 供外部访问的，传入名字和对应的算法
	mgr.registerBalancer(name,b)
}
func DoBalance(name string,insts []*Instance) (inst *Instance,err error) {
	// DoBalance接口，根据不同的name调用不同的balance的DoBalance方法
	balancer,ok := mgr.allBalancer[name] // 拿到对应的balance
	if !ok {
		err = fmt.Errorf("Not found %s balancer",name)
		return
	}
	fmt.Printf("use %s balance\n",name)
	inst,err = balancer.DoBalance(insts)  // 使用balance的DoBalance方法
	return
}