package main

import "fmt"

func main() {
    _, num, strs := numbers() //只获取函数返回值的后两个
    fmt.Println(num, strs)
}

//只获取函数返回值的后两个
func numbers() (int, int, string) {
    a, b, c := 1, 2, "str"
    return a, b, c
}
