// 非递归实现
package main
import "fmt"

func BinarySearch(array []int, target int, l, r int) int {
	ltemp := l
	rtemp := r

	for {
		if ltemp > rtemp {
			return -1
		}

		mid := (ltemp + rtemp) / 2
		middleNum := array[mid]

		if middleNum == target {
			return mid
		} else if middleNum > target {
			rtemp = mid - 1
		} else {
			ltemp = mid + 1
		}
	}
}

func main() {
    array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
    target := 500
    result := BinarySearch(array, target, 0, len(array)-1)
    fmt.Println(target, result)

    target = 189
    result = BinarySearch(array, target, 0, len(array)-1)
    fmt.Println(target, result)
}
