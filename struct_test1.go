package main

import "fmt"
import "unsafe"

func main() {
    // 字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,
    // 这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
    s := "hello"
    fmt.Println(unsafe.Sizeof(s))

    // 在定义常量组时，如果不提供初始值，则表示将使用上行的表达式。
    const (
        a = 1
        b
        c
        d
    )

    fmt.Println(a)
    // b、c、d没有初始化，使用上一行(即a)的值
    fmt.Println(b)   // 输出1
    fmt.Println(c)   // 输出1
    fmt.Println(d)   // 输出1
}

