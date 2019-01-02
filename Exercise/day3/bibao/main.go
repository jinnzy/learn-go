package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	// `json:"username"` 是tag，使用json序列化后会变成username
	UserName string `json:"username"`
	NickName string
	Age int
	Sex string
	Email string
	Phone string
	// 大写的原因是，json在另一个包
}
func testStruct() {
	// 定义数值，后续可以从数据库拿到值
	user1 := &User{
		UserName:"user1",
		NickName:"昵称",
		Age:18,
		Sex:"男",
		Email:"xxx@qq.com",
		Phone:"123123",
	}
	// data取到的数据是byts类型
	data,err := json.Marshal(user1)
	if err != nil{
		fmt.Printf("json.marshal failed,err:",err)
		return
	}
	fmt.Printf("%s\n",string(data))
}
func testMap()  {
	// 定义map
	var m map[string]interface{}
	// 初始化map
	m = make(map[string]interface{})
	m["username"] = "map1"
	m["age"] = 18
	m["sex"] = "man"
	data,err := json.Marshal(m)
	if err != nil{
		fmt.Printf("json err:",err)
	}
	fmt.Printf("%s\n",string(data))
}
func testSlice()  {
	// 定义slice类型，里面每个类型是map，[]后面是类型
	var s []map[string]interface{}
	// 初始化map
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "map1"
	m["age"] = 18
	m["sex"] = "man"
	// 把map追加到slice中
	s = append(s,m)
	m = make(map[string]interface{})
	m["username"] = "map2"
	m["age"] = 123
	m["sex"] = "man123"
	s = append(s,m)
	data,err := json.Marshal(s)
	if err != nil{
		fmt.Printf("json err:",err)
	}
	fmt.Printf("%s\n",string(data))
}
func main()  {
	//testStruct()
	//testMap()
	testSlice()
}