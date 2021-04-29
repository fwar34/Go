package main

import (
	"fmt"
	"sync"
	"time"
)

type Money struct {
	lock sync.Mutex
	amount int64
}

func (m *Money) Add(i int64) {
	m.lock.Lock()
	defer m.lock.Unlock() // 在函数结束后执行

	m.amount = m.amount + i
}

func (m *Money) Minute(i int64) {
	m.lock.Lock()
	defer m.lock.Unlock() // 在函数结束后执行

	if m.amount >= i {
		m.amount = m.amount - i
	}
}

func (m *Money) Get() int64 {
	return m.amount
}

func main() {
    m := new(Money)
	m.Add(10000)

	for i := 0; i < 1000; i++ {
		go func() {
			time.Sleep(500 * time.Millisecond)
			m.Minute(5)
		} ()
	}

	time.Sleep(20 * time.Second)
	fmt.Println(m.Get())
}
