package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Data string
	Left *TreeNode
	Right *TreeNode
}

// 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印根节点
	fmt.Println(tree.Data, "")
	// 再打印左子树
	PreOrder(tree.Left)
	// 再打印右子树
	PreOrder(tree.Right)
}

// 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印左子树
	MidOrder(tree.Left)
	// 再打印根节点
	fmt.Println(tree.Data, " ")
	// 再打印右子树
	MidOrder(tree.Right)
}

// 后续遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印左子树
	PostOrder(tree.Left)
	// 再打印右子树
	PostOrder(tree.Right)
	// 再打印根节点
	fmt.Println(tree.Data, " ")
}

// 层次遍历，用到一种名叫广度遍历的方法，需要使用辅助的先进先出的队列
// 1. 先将树的根节点放入队列。
// 2. 从队列里面 remove 出节点，先打印节点值，如果该节点有左子树节点，左子树入栈，如果有右子树节点，右子树入栈。
// 3. 重复2，直到队列里面没有元素。
func LayerOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 新建队列
	queue := new(LinkQueue)
	// 根节点先入队
	queue.Add(tree)
	for queue.size > 0 {
		// 不断出队
		element := queue.Remove()

		// 先打印节点值
		fmt.Print(element.Data, " ")

		// 左子树非空，入队列
		if element.Left != nil {
			queue.Add(element.Left)
		}

		// 右子树非空，入队列
		if element.Right != nil {
			queue.Add(element.Right)
		}
	}
}

// 链表节点
type LinkNode struct {
	Next *LinkNode
	Value *TreeNode
}

type LinkQueue struct {
	root *LinkNode
	size int
	sync.Mutex
}

// 入队
func (queue *LinkQueue) Add(v *TreeNode) {
	queue.Lock()
	defer queue.Unlock()

	// 如果队列为空，那么增加节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		// 否则新元素插入链表的尾部
		newNode := new(LinkNode)
		newNode.Value = v

		// 一直遍历到链表尾部
		node := queue.root
		for node.Next != nil {
			node = node.Next
		}

		node.Next = newNode
	}

	queue.size = queue.size + 1
}

// 出队
func (queue *LinkQueue) Remove() *TreeNode {
	queue.Lock()
	defer queue.Unlock()

	if queue.size == 0 {
		panic("over limit")
	}

	// 顶部元素要出队
	topNode := queue.root
	v := topNode.Value

	// 将顶部元素的后继链接上
	queue.root = topNode.Next

	queue.size = queue.size - 1
	return v
}

func (queue *LinkQueue) Size() int {
	return queue.size
}

func main() {
	t := &TreeNode {Data: "A"}
	t.Left = &TreeNode {Data: "B"}
	t.Right = &TreeNode {Data: "C"}
	t.Left.Left = &TreeNode {Data: "D"}
	t.Left.Right = &TreeNode {Data: "E"}
	t.Right.Left = &TreeNode {Data: "F"}

	fmt.Println("先序排序：")
	PreOrder(t)
	fmt.Println("\n中序排序：")
	MidOrder(t)
	fmt.Println("\n后序排序")
	PostOrder(t)
	fmt.Println("\n层次排序")
	LayerOrder(t)
}

