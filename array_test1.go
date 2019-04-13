package main

import "fmt"

func main() {
    var balance[10] float32
    fmt.Println(balance)

    var balance2 = [5] float32 {1000.0, 2.0, 3.4, 7.0, 50.0}
    fmt.Println(balance2)

    var balance3 = [5] float32 {1000.0, 2.0, 3.4, 7.0}
    fmt.Println(balance3)

    var balance4 = [...] float32 {1000.0, 2.0, 3.4, 7.0, 50.0}
    fmt.Println(balance4)

    a := [3][4]int{
        {0, 1, 2, 3} ,   /*  第一行索引为 0 */
        {4, 5, 6, 7} ,   /*  第二行索引为 1 */
        {8, 9, 10, 11},   /* 第三行索引为 2 */
    }
    //注意：以上代码中倒数第二行的
    //} 必须要有逗号，因为最后一行的 } 不能单独一行，也可以写成这样
    b := [3][4]int{
        {0, 1, 2, 3} ,   /*  第一行索引为 0 */
        {4, 5, 6, 7} ,   /*  第二行索引为 1 */
        {8, 9, 10, 11}}   /* 第三行索引为 2 */
    fmt.Println(a)
    fmt.Println(b)
}
