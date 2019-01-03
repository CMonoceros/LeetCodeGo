package main

// 4ms
func twoSum(nums []int, target int) []int {
	var m = make(map[int][]int)
	var res []int
	for i, num := range nums {
		v, ok := m[num]
		if ok {
			return append(v, i)
		} else {
			m[target-num] = []int{i}
		}
	}
	return res
}

// 80ms
func twoSum2(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
