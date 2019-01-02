package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score float32
}
func (s Student) Print() {
	fmt.Println("-------------")
	fmt.Println(s)
	fmt.Println("-------------")
}
//func (s *Student) Set(name string,age int,score float32)  { // 这个传入指针的话，下面获取method的数量会是0，找不到方法
func (s Student) Set(name string,age int,score float32)  {
	s.Name = name
	s.Age = age
	s.Score = score
}

func TestStruct(a interface{})  {
	val := reflect.ValueOf(a) //获取值，这个类型是reflect.Value类型
	kd := val.Kind() // 打印底层类型
	// 判断类型reflect这个包里有所有的常亮，可以用于判断
	if kd != reflect.Struct {
		// 如果不是结构体，则退出
		fmt.Println("不是struct")
		return
	}
	num := val.NumField() // 可以获得字段的数量
	fmt.Printf("struct有%d个字段\n", num)
	fmt.Println(val.Field(0)) // 取结构体的0位置值，输出是stu1

	numOFMethod := val.NumMethod() // 输出方法数量
	fmt.Printf("struct有%d个方法\n", numOFMethod)

	// 定义call传入方法的值
	var params []reflect.Value
	val.Method(0).Call(params) // 调用
}

func main()  {
	var a Student = Student{
		Name: "stu1",
		Age: 18,
		Score: 92.8,
	}
	TestStruct(a) // 传的值类型，如果要传指针类型，方法那里要加*和val.NumMethod要加Elem()
}
