// 开始有10000,没有给 money 加锁，所以最后转账1000次每次5元后余额不是5000
package main

import (
	"fmt"
	"time"
)

type Money struct {
	amount int64
}

// 加钱
func (m *Money) Add(i int64) {
	m.amount = m.amount + i
}

// 减钱
func (m *Money) Minute(i int64) {
	// 钱足才能减
	if m.amount >= i {
		m.amount = m.amount - i
	}
}

// 查看还有多少钱
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

