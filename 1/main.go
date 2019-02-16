package main

import "fmt"

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

func twoSum(nums []int, target int) []int {
	hash_map := make(map[int]int)
	var a []int

	for k, n := range nums {

		if _, ok := hash_map[target-n]; ok {
			fmt.Println(k, hash_map[target-n])

			a = []int{k, hash_map[target-n]}
			return a
		}
		hash_map[n] = k
	}
	return a
}

func main() {
	var nums []int

	nums = []int{2, 7, 11, 15}

	target := 9

	b := twoSum(nums, target)
	fmt.Println(b)
}
