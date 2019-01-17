package test

import "encoding/json"

type student struct {
	Name string
	Sex string
	Age int
}

func (p *student) Save()  {
	data,err := json.Marshal(p)
}
