package model

import "github.com/pkg/errors"

type Student struct {
	Name string
	Grade string
	Id string
	Sex string
	books []*BorrowTiem // 列表，值是存的指针，用[]Book占用内存多，每个还要拷贝下
}
// 借书和借书的数量的结构体
type BorrowTiem struct {
	book *Book
	num int
}
// 使用工厂模式，传入值并返回
func CreateStudent(name,grade,id,sex string) *Student {
	stu := &Student{
		Name: name,
		Grade:grade,
		Id:id,
		Sex:sex,
	}
	return stu
}
// 增加借书方法
func (s *Student) AddBook(b *BorrowTiem) {
	// 传入b 书籍，在原有基础上加入b书籍
	s.books = append(s.books,b)
}

func (s *Student) DelBook(b *BorrowTiem) (err error) {
	for i := 0; i < len(s.books); i++ {
		if s.books[i].book.Name == b.book.Name {
			if b.num == s.books[i].num {
				// 切片到i的位置，不包含i
				front := s.books[0:i]
				// 从i+1切片包含i，去尾留头，正好把i剔除了
				left := s.books[i+1:]
				front = append(front, left...)
				s.books = front
			}
			s.books[i].num -= b.num
		}
	}
	err = errors.New("书籍错误")
	return err
}