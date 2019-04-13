package main

import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main() {
    fmt.Println(Books{"Go language", "www.runoob.com", "Go language book", 343434})
    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
    // 忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
