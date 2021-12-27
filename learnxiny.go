package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []int32{1, 2, 3}
	fmt.Println(s)
	fmt.Println(append(s, []int32{4, 5, 6}...))
	x, y := memoryTest()
	fmt.Println(*x, *y)
	*y = 4
	fmt.Println(*x, *y)
	varTest("aaa", "bbb", "ccc")
	errTest()
	conCurrency()
}

func memoryTest() (x, y *int) {
	x = new(int)
	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r
}

func varTest(myStrings ...interface{}) {
	for i, v := range myStrings {
		fmt.Println(i, v)
	}
}

func errTest() {
	m := map[int]string{
		3: "three", 4: "four",
	}

	if x, ok := m[2]; !ok {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found:", x)
	}

	if _, err := strconv.Atoi("nonlsls"); err != nil {
		fmt.Println(err)
	}
}

func inc(i int, c chan int) {
	c <- i + 1
}

func conCurrency() {
	c := make(chan int)
	go inc(0, c)
	go inc(10, c)
	go inc(-805, c)
	fmt.Println(<-c, <-c, <-c)
}
