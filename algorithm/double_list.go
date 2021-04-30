package main

import (
	"fmt"
	"sync"
)

type DoubleList struct {
	head *ListNode // 链表头部
	tail *ListNode // 链表尾部
	size int // 链表长度
	lock sync.Mutex
}

type ListNode struct {
	prev *ListNode // 前驱
	next *ListNode // 后驱
	value string // 值
}

// 获取节点值
func (node *ListNode) Get() string {
	return node.value
}

// 获取节点的前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.prev
}

// 获取节点的后驱节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// 是否存在后驱节点
func (node *ListNode) HasNext() bool {
	return node.next != nil
}

// 是否存在前驱节点
func (node *ListNode) HasPre() bool {
	return node.prev != nil
}

// 节点是否为空
func (node *ListNode) IsNil() bool {
    return node == nil
}

// 添加节点到链表的第 N 个元素之前，N 为0表示新节点成为链表的头部
func (list *DoubleList) AddNodeFromHead(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超过链表长度
	if n > list.len {
		panic("index out")
	}

	// 先找出头部
	node := list.head

	// 往后遍历拿到的 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 新节点
	newNode := new(ListNode)
	newNode.value = v

	// 如果定位到节点为空，表示链表为空，将新节点设置为新头部和新尾部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到的节点，它的前驱
		prev := node.prev

		// 如果定位到的节点前驱为 nil, 那么定位到的节点为链表头部，需要换头部
		if prev.IsNil() {
			// 将新节点放在头部
			newNode.next = node
			node.prev = newNode
			//新节点成为头部
			list.head = newNode
		} else {
			// 将新节点插入到定位到的节点之前
			// 定位到的节点的前驱节点 prev 现在链接到新节点上
			prev.next = newNode
			newNode.prev = prev
			// 定位到的节点现在链接到新节点上
			node.prev = newNode
			newNode.next = node
		}
	}
	// 链表长度加1
	list.size = list.size + 1
}

// 添加节点到链表尾部的第 N 个元素之后， N=0 表示新节点成为新的尾部
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()

	if n > list.size {
		panic("index out")
	}

	node := list.tail

	// 往前遍历拿到第 N+1 个元素
	for i := 1; i <= n; i++ {
		node = node.prev
	}

	newNode := new(ListNode)
	newNode.value = v

	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到节点，它的后驱
		next := node.next

		if next.IsNil() {
			node.next = newNode
			newNode.prev = node

			// 新节点成为尾部
			list.tail = newNode
		} else {
			newNode.prev = node
			node.next = newNode
			newNode.next = next
			next.prev = newNode
		}
	}

	list.size = list.size + 1
}

// 返回列表链表头结点
func (list *DoubleList) First() *ListNode {
    return list.head
}

// 返回列表链表尾结点
func (list *DoubleList) Last() *ListNode {
    return list.tail
}

// 从头部开始往后找，获取第N+1个位置的节点，索引从0开始。
func (list *DoubleList) IndexFromHead(n int) *ListNode {
    // 索引超过或等于列表长度，一定找不到，返回空指针
    if n >= list.len {
        return nil
    }

    // 获取头部节点
    node := list.head

    // 往后遍历拿到第 N+1 个位置的元素
    for i := 1; i <= n; i++ {
        node = node.next
    }

    return node
}

// 从尾部开始往前找，获取第N+1个位置的节点，索引从0开始。
func (list *DoubleList) IndexFromTail(n int) *ListNode {
    // 索引超过或等于列表长度，一定找不到，返回空指针
    if n >= list.len {
        return nil
    }

    // 获取尾部节点
    node := list.tail

    // 往前遍历拿到第 N+1 个位置的元素
    for i := 1; i <= n; i++ {
        node = node.pre
    }

    return node
}
