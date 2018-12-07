package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name string
	Age int		// Age大写是公开的包外可以访问，和函数的一样
	score float32
	next *Student // 值默认是nil ，指针类型，留着存指向下一个链表的地址
}
func travers(p *Student)  {
	//var p *Student = &head // 定义一个指针变量类型是Student的，指向的是head头的地址
	// 循环放到函数里了 所以上面注释掉
	for p != nil {
		fmt.Println(*p)
		p = p.next // 和 (*p).next写法一样，指向下一个节点，继续循环
	}
}
func insertTail(p *Student)  {
	// 先传进来头部地址
	tail := p // tail指向头部地址
	for i := 0;i < 10; i++ {
		// 循环创建结构体节点
		stu := Student{
			Name: fmt.Sprintf("stu%d", i),
			Age: rand.Intn(100),
			score: rand.Float32() * 100,
		}
		// 创建完节点之后，把tauk.next地址指向刚创建的结构体节点stu
		tail.next = &stu
		tail = &stu // 然后再把stu地址传给tail，tail就是最后一个节点，继续循环，会一直把新生成的结构体节点加入到最后。
	}
}
func main()  {
	var head Student
	head.Name = "hua"
	head.Age = 18
	head.score= 100

	insertTail(&head)
	travers(&head)
}
