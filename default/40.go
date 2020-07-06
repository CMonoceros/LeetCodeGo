package main

import "sort"

// Array
// Backtracking
// 12ms
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var now []int
	sort.Ints(candidates)

	var dfs func(candidates []int, target, index, sum int, res *[][]int, now []int)
	dfs = func(candidates []int, target, index, sum int, res *[][]int, now []int) {
		if sum == target {
			*res = append(*res, now)
			return
		}
		if sum > target || index >= len(candidates) {
			return
		}
		for i := index + 1; i < len(candidates); i++ {
			cp := make([]int, len(now)+1)
			copy(cp, append(now, candidates[i]))
			dfs(candidates, target, i, sum+candidates[i], res, cp)
			for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
				i++
			}
		}
	}

	dfs(candidates, target, 0, 0, &res, now)
	return res
}
