package main

import "fmt"

func main() {
	var a = []int{0, 2, 1, 0}
	r := peakIndexInMountainArray(a)
	fmt.Println(r)
}

func peakIndexInMountainArray(A []int) int {
	var tmp int
	for k, _ := range A {
		if A[k] > A[k+1] {
			tmp = k
			break
		}

	}

	return tmp
}
