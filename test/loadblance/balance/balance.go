package balance

type Balancer interface {
	// 定义接口blance，实现不同的负载均衡算法，传进去实例列表，返回实例
	DoBalance([]*Instance, ...string) (*Instance,error)
	// ...string 是合并参数，可变参数，实现的时候要加上，但是用不用都可以
}