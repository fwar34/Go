// 二分查找  递归
package main
import "fmt"

func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		// 出界了
		return -1
	}

	// 从中间开始查找
	mid := (l + r) / 2
	middleNum := array[mid]

	if middleNum == target {
		return mid
	} else if middleNum > target {
		return BinarySearch(array, target, 0, mid - 1)
	} else {
		return BinarySearch(array, target, mid + 1, r)
	}
}

func main() {
    array := []int{1, 5, 9, 15, 19, 123, 189, 200}
	target := 123
	result := BinarySearch(array, target, 0, len(array) - 1)
	fmt.Println(target, result)

	target = 500
	result = BinarySearch(array, target, 0, len(array) - 1)
	fmt.Println(target, result)
}

