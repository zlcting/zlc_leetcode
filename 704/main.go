package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 5}
	target := 2
	a := search(nums, target)

	fmt.Println(a)
}

func search(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for {
		// mid向下取整
		mid = ((left + right) / 2)
		if nums[mid] > target { // 如果当前元素大于k，那么把right指针移到mid - 1的位置
			right = mid - 1
		} else if nums[mid] < target { // 如果当前元素小于k，那么把left指针移到mid + 1的位置
			left = mid + 1
		} else { // 否则就是相等了，退出循环
			break
		}

		if left > right { // 判断如果left大于right，那么这个元素是不存在的。返回-1并且退出循环
			mid = -1
			break
		}
	}
	return mid
}
