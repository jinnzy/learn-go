package blance

type blance interface {
	// 定义接口blance，实现不同的负载均衡算法，传进去实例列表，返回实例
	DoBalance([]*Instance) (*Instance,error)
}