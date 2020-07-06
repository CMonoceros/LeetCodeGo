package main

import "math"

// Backtracking

// 0ms
func totalNQueens(n int) int {
	res := 0
	var dfs func(loc []int, index int)
	var checkValid func(loc []int, index int) bool
	dfs = func(loc []int, index int) {
		if index == n {
			res++
		} else {
			for i := 0; i < n; i++ {
				loc[index] = i
				if checkValid(loc, index) {
					dfs(loc, index+1)
				}
			}
		}
	}
	checkValid = func(loc []int, index int) bool {
		for i := 0; i < index; i++ {
			if loc[index] == loc[i] {
				return false
			}
			if int(math.Abs(float64(loc[index]-loc[i]))) == index-i {
				return false
			}
		}
		return true
	}

	dfs(make([]int, n), 0)
	return res
}
