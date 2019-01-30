package main

import "math"

// Backtracking

// todo:
// 12ms
func solveNQueens(n int) [][]string {
	var res [][]string
	var dfs func(loc []int, index int)
	var checkValid func(loc []int, index int) bool
	dfs = func(loc []int, index int) {
		if index == n {
			var now []string
			for i := 0; i < n; i++ {
				s := ""
				for j := 0; j < n; j++ {
					if j == loc[i] {
						s += "Q"
					} else {
						s += "."
					}
				}
				now = append(now, s)
			}
			res = append(res, now)
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
