package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score int
}

func (p *Student) init(name string,age int,score int) {
	// 要接收指针类型，下面的修改才会生效
	// p代表当前这个struct的实例、
	// 在Student类型中定义一个这个方法
	p.Name = name
	p.Age = age
	p.Score = score
	fmt.Println(p)
}
func (p Student) get() Student  {
	// Student 是返回值
	return p
}
func main() {
	var stu Student // 初始化类型，这里p就等于stu
	//(&stu).init("stu",10,200) // 调用类型中的方法，传入地址
	stu.init("stu",10,200)  //上面那样写很麻烦，go会自动帮你转成指针
	// stu.init 传入的和函数一样是一个副本，要传入地址才行，用(&stu)
	stu1 := stu.get()
	fmt.Println(stu1)
}