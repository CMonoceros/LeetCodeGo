package main

import (
	"container/list"
	"math"
)

// Array
// Two Pointers
// Stack
// Dynamic Programming

// todo: 1
// 76ms
func trap(height []int) int {
	res := 0
	for i := 1; i < len(height); i++ {
		lMax, rMax := 0, 0
		for j := i; j >= 0; j-- {
			if lMax < height[j] {
				lMax = height[j]
			}
		}
		for j := i; j < len(height); j++ {
			if rMax < height[j] {
				rMax = height[j]
			}
		}
		res += int(math.Min(float64(lMax), float64(rMax))) - height[i]
	}
	return res
}

// DP缓存查询结果
// todo: 1
// 4ms
func trap2(height []int) int {
	res := 0
	lMax := make([]int, len(height))
	rMax := make([]int, len(height))
	max := 0
	for i := 0; i < len(height); i++ {
		if height[i] >= max {
			max = height[i]
		}
		lMax[i] = max
	}
	max = 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] >= max {
			max = height[i]
		}
		rMax[i] = max
	}
	for i := 1; i < len(height); i++ {
		res += int(math.Min(float64(lMax[i]), float64(rMax[i]))) - height[i]
	}
	return res
}

// 结果区间总会从两指针的较小值中所取
// todo:
// 4ms
func trap3(height []int) int {
	res := 0
	lMax, rMax := 0, 0
	l, r := 0, len(height)-1
	for l < r {
		if height[l] < height[r] {
			if height[l] >= lMax {
				lMax = height[l]
			} else {
				res += lMax - height[l]
			}
			l++
		} else {
			if height[r] >= rMax {
				rMax = height[r]
			} else {
				res += rMax - height[r]
			}
			r--
		}
	}
	return res
}

// 栈内数据递减，遇到较大值时，出栈并计算两点间距离
// todo:
// 4ms
func trap4(height []int) int {
	res := 0
	stack := list.New()
	for i := 0; i < len(height); i++ {
		for stack.Len() != 0 && height[i] > height[stack.Back().Value.(int)] {
			top := stack.Back()
			stack.Remove(top)
			if stack.Len() == 0 {
				break
			}
			l := stack.Back().Value.(int)
			width := i - l - 1
			height := int(math.Min(float64(height[i]), float64(height[l]))) - height[top.Value.(int)]
			res += height * width
		}
		stack.PushBack(i)
	}
	return res
}
