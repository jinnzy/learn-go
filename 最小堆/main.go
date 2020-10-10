package main

import "fmt"

type Heap []int

func (h Heap) swap(i, j int) {

	h[i], h[j] = h[j], h[i]

}

func (h Heap) less(i, j int) bool {

	return h[i] < h[j]

}

func (h Heap) up(i int) {

	for {

		f := (i - 1) / 2 // 父亲结点

		if i == f || h.less(f, i) {

			break

		}

		h.swap(f, i)

		i = f

	}

}

func (h *Heap) Push(x int) {

	*h = append(*h, x)

	h.up(len(*h) - 1)

}

func (h Heap) down(i int) {

	for {

		l := 2*i + 1 // 左孩子

		if l >= len(h) {

			break // i已经是叶子结点了

		}

		j := l // j为左孩子的位置
		// r为右孩子的位置，如果小于数组长度则存在，并且右孩子比左孩子小，则替换位置
		if r := l + 1; r < len(h) && h.less(r, l) {
			// j替换为右孩子位置
			j = r // 右孩子

		}
		// 比对i和j的位置，i是j的父节点
		if h.less(i, j) {

			break // 如果父结点比孩子结点小，则不交换

		}

		h.swap(i, j) // 交换父结点和子结点

		i = j        //继续向下比较

	}

}

// 删除堆中位置为i的元素

// 返回被删元素的值

func (h *Heap) Remove(i int) (int, bool) {

	if i < 0 || i > len(*h)-1 {

		return 0, false

	}

	n := len(*h) - 1

	h.swap(i, n) // 用最后的元素值替换被删除元素

	// 删除最后的元素

	x := (*h)[n]

	*h = (*h)[0:n]

	// 如果当前元素大于父结点，向下筛选

	if (*h)[i] > (*h)[(i-1)/2] {

		h.down(i)

	} else { // 当前元素小于父结点，向上筛选

		h.up(i)

	}

	return x, true

}


// 弹出堆顶的元素，并返回其值

func (h *Heap) Pop() int {

	n := len(*h) - 1
	// 要删除0位置的值，想将0的值和最后的值交换
	h.swap(0, n)
	x := (*h)[n]
	// 删除n
	*h = (*h)[0:n]
	// 将0值进行下沉
	h.down(0)

	return x

}

func (h Heap) Init() {
	n := len(h)

	// i > n/2-1 的结点为叶子结点本身已经是堆了

	for i := n/2 - 1; i >= 0; i-- {

		h.down(i)

	}

}


func main() {

	var h = Heap{20, 7, 3, 10, 15, 25, 30, 17, 19}

	h.Init()

	fmt.Println(h) // [3 7 20 10 15 25 30 17 19]

	h.Push(6)

	fmt.Println(h) // [3 6 20 10 7 25 30 17 19 15]

	x, ok := h.Remove(5)

	fmt.Println(x, ok, h) // 25 true [3 6 15 10 7 20 30 17 19]

	y, ok := h.Remove(1)

	fmt.Println(y, ok, h) // 6 true [3 7 15 10 19 20 30 17]

	z := h.Pop()

	fmt.Println(z, h) // 3 [7 10 15 17 19 20 30]

}
