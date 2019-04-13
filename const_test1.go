package main

import "fmt"

func main() {
    const LENGTH int = 8
    const WIDTH int = 5
    var area int
    const a, b, c = 1, false, "str"

    area = LENGTH * WIDTH
    fmt.Printf("area : %d", area)
    println()
    println(a, b, c)
}

// 常量还可以用作枚举： 
const (
    Unknow = 0
    Female = 1
    Male = 2
)
