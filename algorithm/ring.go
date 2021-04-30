// https://goalgorithm.github.io/algorithm/link.html
package main
import "fmt"

type Ring struct {
	next, prev *Ring // 前驱和后驱节点
	Value interface {} // 数据
}

// 初始化循环链表
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// 创建 N 个节点的循环链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}

	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// 获取第 N 个节点， n 为负数往前面遍历
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}

	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}

	return r
}

// 往节点 A, 链接一个节点(这个节点有可能是一个新节点或者是另一个链表)，并且返回之前节点 A 的后驱节点
// 那么也就是在 r 节点后插入一个新节点 s，而 r 节点之前的后驱节点，将会链接到新节点后面，并返回 r 节点之前的第一个后驱节点 n
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 删除节点后面的 n 个节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}

	return r.Link(r.Move(n + 1))
}

func linkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接5个新节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	node := r
	for {
		fmt.Println(node.Value)

		// 移动到下一节点
		node = node.Next()

		// 如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}

// 获取链表长度
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.Next() {
			n++
		}
	}
	return n
}

func deleteTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接4个新节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	temp := r.Unlink(3) // 解除了后面3个节点

	fmt.Println("deleteTest")
	// 打印原来的节点
	node := r
	for {
		fmt.Println(node.Value)
		node = node.Next()

		if node == r {
			break
		}
	}

	fmt.Println("--------------------")

	// 打印被切断的节点
	node = temp
	for {
		fmt.Println(node.Value)
		node = node.Next()

		if node == temp {
			break
		}
	}
}

func main() {
	linkNewTest()
	deleteTest()
}
