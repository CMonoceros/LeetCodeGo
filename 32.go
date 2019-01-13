package main

import "math"

// String
// Dynamic Programming

// 36ms
func longestValidParentheses(s string) int {
	re := []rune(s)
	res := 0
	d := make([]p, len(re))
	for i := 0; i < len(re)-1; i++ {
		for j, l, r := len(re)-1-i, 0, 0; j >= 0; j-- {
			if string(re[j]) == ")" {
				r++
			} else if string(re[j]) == "(" {
				l++
			}
			if l > r {
				break
			}
			if l == r {
				res = int(math.Max(float64(res), float64(len(re)-i-j)))
			}
			d[j] = p{
				l: l,
				r: r,
			}
		}
	}
	return res
}

type p struct {
	l int
	r int
}

// todo:
// 4ms
func longestValidParentheses2(s string) int {
	re := []rune(s)
	if len(re) == 0 {
		return 0
	}
	d := make([]int, len(re))
	d[0] = 0
	res := 0
	for i := 1; i < len(re); i++ {
		if string(re[i]) == ")" && string(re[i-1]) == "(" {
			d[i] = d[i-1] + 2
			if i-2 >= 0 {
				d[i] += d[i-2]
			}
		} else if string(re[i]) == ")" && string(re[i-1]) == ")" &&
			i-1-d[i-1] >= 0 && string(re[i-1-d[i-1]]) == "(" {
			d[i] = d[i-1] + 2
			if i-1-d[i-1]-1 >= 0 {
				d[i] += d[i-1-d[i-1]-1]
			}
		} else {
			d[i] = 0
		}
		res = int(math.Max(float64(res), float64(d[i])))
	}
	return res
}
