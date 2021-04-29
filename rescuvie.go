// 尾递归
package main
import "fmt"

func F(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}

	return F(n - 1, a2, a1 + a2)
}

func main() {
    fmt.Println(F(5, 1, 1))
}

