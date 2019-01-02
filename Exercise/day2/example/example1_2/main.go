package main

import (
	"fmt"
	"reflect"
)

func dn() {
	var m = make(map[string]int)
	m["sdf"] = 1
	m["sdf1"] = 1

	type sdra map[string]int
	//type sdra  map[string]interface{}

	var dragon = make(map[string]interface{})
	dragon["abs"] = 1
	dragon["absd"] = nil
	dragon["asbd"] = m

//for v, s := range dragon {
//if t, ok := s.(sdra); ok {
//fmt.Println("s是sdra类型:", t)
//}
//fmt.Println(v, reflect.TypeOf(s))
//a := reflect.ValueOf(v)
//fmt.Println("kind:",v, a.Kind())
//}
	for v, s := range dragon {
		if t, ok := s.(sdra); ok {
			fmt.Println("s是sdra类型:", t)
		}
		fmt.Println(v, reflect.TypeOf(s))
	}
}


func main()  {
	dn()
}