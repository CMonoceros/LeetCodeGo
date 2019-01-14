package main

// Array
// Divide and Conquer
// Dynamic Programming

// 8ms
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	now := nums[0]
	for i := 1; i < len(nums); i++ {
		if now < 0 {
			now = nums[i]
		} else {
			now += nums[i]
		}
		if now > res {
			res = now
		}
	}
	return res
}
