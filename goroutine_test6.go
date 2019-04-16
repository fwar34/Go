package main

import "fmt"

func log(s string, c chan int) {
	fmt.Println("Hello", s)
	c <- 0
}

func main() {
	ch := make(chan int)
	go log("jing", ch)
	<-ch
}
