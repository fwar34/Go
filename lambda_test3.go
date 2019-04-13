package main

import "fmt"

func main() {
    add_func := add(1, 2)
    fmt.Println(add_func(1, 1))
    fmt.Println(add_func(2, 2))
    fmt.Println(add_func(3, 3))
}

func add(a, b int) func(x1 int, x2 int) (int, int, int) {
    i := 0
    return func(x1 int, x2 int) (int, int, int) {
        i++
        return i, a + b, x1 + x2
    }
}
