package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limitor := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limitor
		fmt.Println("request", req, time.Now())
	}

	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	burstyLimitor := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimitor <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimitor <- t
		}
	}()

	for request := range burstyRequests {
		<-burstyLimitor
		fmt.Println("request", request, time.Now())
	}
}
