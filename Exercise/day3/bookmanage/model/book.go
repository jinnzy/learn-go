package model

import "github.com/pkg/errors"


type Book struct {
	Name string
	Total int
	Author string
	CreateTime string
}

// 使用工厂模式，传入Book所需的值,返回b
func CreateBook(name string, total int,author string,createtime string)(b *Book) {
	b = &Book{
		Name:name,
		Total:total,
		Author:author,
		CreateTime:createtime,
	}
	return b
}
func (b *Book) CanBorrow(c int) bool {
	return b.Total >= 0
}
// 借书
func (b *Book) Borrow(c int) (err error) {
	if b.CanBorrow(c) == false {
		err = errors.New("库存不足")
		return
	}
	b.Total -= c
	return
}
// 还书
func (b *Book) Back(c int)  {
	b.Total += c
	return
}