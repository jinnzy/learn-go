package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//var nodeTmp *ListNode
	var node = &ListNode{0,nil} // 预先赋值，是为了后面的nodeTmp = node 会被指到node的内存地址
	var carry int
	var sum int
	nodeTmp := node // 后面操作都用nodeTmp操作，实际返回node的值，是因为nodeTmp用next往下走
	// l1 l2 等于的情况时，carry 大于0 进位了，还要继续next
	for l1 != nil || l2 != nil || carry >0 {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if carry > 0 {
			sum += carry
			carry = 0
		}
		if sum >= 10 {
			sum = sum % 10 // 求整除后的数 15 / 10 = 1
			carry = 1
		}
		if nodeTmp == nil {
			nodeTmp = &ListNode{sum,nil}
		}else {
			nodeTmp.Next = &ListNode{sum,nil}
			nodeTmp = nodeTmp.Next
		}
	}
	return node.Next
}
type ListNode struct {
	     Val int
	     Next *ListNode
	 }
//func main()  {
//	//var a1 *ListNode
//	var l1 *ListNode
//	l1 = &ListNode{0,nil}
//	 a1 := l1
//	//a1 = &ListNode{0,nil}
//	a1.Next= &ListNode{2,nil}
//	a1 = a1.Next
//	a1.Next = &ListNode{4,nil}
//	a1 = a1.Next
//	a1.Next = &ListNode{3,nil}
//
//
//	//fmt.Println(a)
//	fmt.Println(l1.Next)
//	fmt.Println(l1.Next.Next)
//	//fmt.Println(a1)
//}
