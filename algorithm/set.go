package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type Set struct {
	m map[int]struct{} // 用字典来实现，因为字段键不能重复
	len int
	sync.RWMutex
}

// 新建一个空集合
// 使用一个容量为 cap 的 map 来实现不可重复集合。map 的值我们不使用，所以值定义为空结构体 struct{}，因为空结构体不占用内存空间
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set {
		m: temp,
	}
}

// 添加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{} // 添加的 value 是一个空结构
	s.len = len(s.m) // 重新计算元素数量
}

// 删除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()

	if s.len == 0 {
		return
	}

	delete(s.m, item) // 实际从 map 删除这个元素
	s.len = len(s.m)
}

// 重看元素是否在集合中
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 查看集合大小
func (s *Set) Len() int {
	return s.len
}

func (s *Set) IsEmpty() bool {
	return s.len == 0
}

// 清除所有集合
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()

	s.m = map[int]struct{}{} // 字典重新赋值
	s.len = 0
}

// 将集合转换成 list
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// 为什么使用空结构体
func other() {
	a := struct{}{}
	b := struct{}{}
	if a == b {
		fmt.Printf("right:%p\n", &a)
	}

	fmt.Println(unsafe.Sizeof(a))
}

func main() {
	other()

	s := NewSet(5)
	s.Add(1)
	s.Add(1)
	s.Add(2)
	fmt.Println("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Has(2) {
		fmt.Println("2 exists")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}

