package main

import (
	"fmt"
	"math/rand"
)

type Test interface {
	Print()
} 

type Student struct {
	name string
	age int
	score int
}


func (p *Student) Print()  {
	fmt.Println("name",p.name)
	fmt.Println("age",p.age)
	fmt.Println("score",p.score)
}

func main()  {
	var t Test // 定义t变量接口
	var stu Student = Student{
		// stu 实现了Test的接口
		name: "stu1",
		age: 12,
		score: 200,
	}
	t = &stu // 可以用t代替stu
	t.Print()
	a := rand.Int31n(1000)
	fmt.Printf("a:%d", a)
}
