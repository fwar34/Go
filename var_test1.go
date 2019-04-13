package main

import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)

    var intVal int
    intVal2 := 1
    intVal, intVal3 := 3, 4
    fmt.Printf("%v %v %v\n", intVal, intVal2, intVal3)
}
