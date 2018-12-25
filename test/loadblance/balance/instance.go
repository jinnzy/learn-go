package balance

type Instance struct {
	host string
	port int

}
// 这里Instance里的host和port都是小写，在main函数中访问不到，所以使用下面的工厂模式（类似构造函数）赋值返回
func NewInstance(host string,port int) *Instance{
	return &Instance{
		host:host,
		port:port,
	}
}

func (p *Instance)GetHost() string  {
	return p.host
}
func (p *Instance)GetPort() int  {
	return p.port
}