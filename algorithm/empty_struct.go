package main

import (
	"fmt"
	"unsafe"
)

// 空结构体的内存地址都一样，并且不占用内存空间
func main() {
	a := struct{}{}
	b := struct{}{}
	if a == b {
		fmt.Printf("right:%p\n", &a)
	}
	fmt.Println(unsafe.Sizeof(a))
}

