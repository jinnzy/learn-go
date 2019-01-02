package balance

import "strconv"

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
// 默认输出是&{192.168.42.59 8080} 这样的，自己实现String方法自定义输出
//结果为192.168.122.254:8080
func (p *Instance)String() string {
	return p.host + ":" + strconv.Itoa(p.port)
}