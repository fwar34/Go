package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// for j := range jobs {
	// 	fmt.Println("worker", id, "started job", j)
	// 	time.Sleep(time.Second)
	// 	fmt.Println("worker", id, "finish job", j)
	// 	results <- j * 2
	// }
	for {
		select {
		case j, more :=  <- jobs:
			if more {
				fmt.Println("worker", id, "started job", j)
				time.Sleep(time.Second)
				fmt.Println("worker", id, "finish job", j)
				results <- j * 2
			} else {
				fmt.Println("no more jobs, worker", id, "stop")
				return
			}
		}
	}
}

func main() {
	const numJobs = 5
    jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 0; a < numJobs; a++ {
		<-results
	}
}

