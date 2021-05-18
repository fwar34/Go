package main

import "fmt"

func SelectSort(list []int) {
	n := len(list)

	for i := 0; i < n - 1; i++ {
		// 每次从第 i 位开始，找到最小的元素
		min := list[i] // 最小数
		minIndex := i // 最小数下标
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				// 如果找到的数比上次的还小，那么最小数变为它
				min = list[j]
				minIndex = j
			}
		}

		// 这一轮找到的最小数下标不等于上次的下标，交换元素
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

// 上面的算法需要从某个数开始，一直扫描到尾部，我们可以优化算法，使得复杂度减少一半。
// 我们每一轮，除了找最小数之外，还找最大数，然后分别和前面和后面的元素交换，这样循环次数减少一半
func SelectSort2(list []int) {
	n := len(list)
	// 只需循环一半
	for i := 0; i < n / 2; i++ {
		minIndex := i
		maxIndex := i

		// 在这一轮迭代中要找到最大值和最小值的下标
		for j := i + 1; j < n - i; j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j // 这一轮是最大的，直接 continue
				continue
			}

			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		if maxIndex == i && minIndex != n - i - 1 {
			// 如果最大值是开头的元素，而最小值不是最尾的元素
            // 先将最大值和最尾的元素交换
			list[n - i - 1], list[maxIndex] = list[maxIndex], list[n - i - 1]
			// 然后最小的元素放在最开头
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == n - i - 1 {
			// 如果最大值在开头，最小值在结尾，直接交换
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			list[i], list[minIndex] = list[minIndex], list[i]
			list[n - i - 1], list[maxIndex] = list[maxIndex], list[n - i - 1]
		}
	}
}

func main() {
    list0 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
    SelectSort(list0)
    fmt.Println(list0)

	list := []int{5}
    SelectSort2(list)
    fmt.Println(list)

    list1 := []int{5, 9}
    SelectSort2(list1)
    fmt.Println(list1)

    list2 := []int{5, 9, 1}
    SelectSort2(list2)
    fmt.Println(list2)

    list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
    SelectSort2(list3)
    fmt.Println(list3)

    list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6}
    SelectSort2(list4)
    fmt.Println(list4)
}
