package main

// Backtracking

// 4ms
func permute(nums []int) [][]int {
	var res [][]int
	visit := make([]bool, len(nums))
	var dfs func(now []int)
	dfs = func(now []int) {
		if len(now) == len(nums) {
			res = append(res, now)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !visit[i] {
				visit[i] = true
				dfs(append(now, nums[i]))
				visit[i] = false
			}
		}
	}
	dfs([]int{})
	return res
}
