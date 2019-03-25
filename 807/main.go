package main

import "fmt"

func main() {

	var grid = [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	a := maxIncreaseKeepingSkyline(grid)
	fmt.Println(a)

}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	max_row := make([]int, len(grid))
	max_col := make([]int, len(grid[0]))
	for m := range grid[0] {
		max_col[m] = grid[0][m]
		for n := range grid {
			if max_col[m] < grid[n][m] {
				max_col[m] = grid[n][m]
			}
		}
	}
	sum := 0

	for i, row := range grid {
		max_row[i] = max(row)
		for j, num := range row {
			min := 0
			if max_col[j] > max_row[i] {
				min = max_row[i]
			} else {
				min = max_col[j]
			}
			sum += min - num
		}
	}

	return sum
}

func max(slice []int) int {
	m := slice[0]
	for _, n := range slice {
		if m < n {
			m = n
		}
	}
	return m
}
