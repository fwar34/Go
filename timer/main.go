package main

import (
	"log"
	"time"

	"github.com/antlabs/timer"
)

func main() {
    tm := timer.NewTimer()

	// 一次性执行，2秒后执行
	tm.AfterFunc(2 * time.Second, func() {
		log.Printf("2 time.Second")
	})

	tk3 := tm.AfterFunc(3 * time.Second, func() {
		log.Printf("3 time.Second")
	})

	tk3.Stop()

	// 周期执行，每一秒执行一次
	tm.ScheduleFunc(1 * time.Second, func() {
		log.Printf("schedule")
	})

	tm.Run()
}

