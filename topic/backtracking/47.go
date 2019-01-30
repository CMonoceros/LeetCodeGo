package main

import "sort"

// Backtracking

// 12ms
func permuteUnique(nums []int) [][]int {
	var res [][]int
	visit := make([]bool, len(nums))
	sort.Ints(nums)
	var dfs func(now []int)
	dfs = func(now []int) {
		if len(now) == len(nums) {
			res = append(res, now)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !visit[i] {
				visit[i] = true
				var cp = make([]int, len(now)+1)
				copy(cp, append(now, nums[i]))
				dfs(cp)
				visit[i] = false
				for ; i+1 < len(nums) && nums[i] == nums[i+1]; {
					i++
				}
			}
		}
	}
	dfs([]int{})
	return res
}
