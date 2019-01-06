package main

// Hash Table
// Two Pointers
// String

import (
	"math"
	"strings"
)

// 496ms
func lengthOfLongestSubstring(s string) int {
	var m = make(map[string]string)
	res := 0
	for _, sub := range s {
		for k, v := range m {
			if strings.Contains(v, string(sub)) {
				res = int(math.Max(float64(len(v)), float64(res)))
				delete(m, k)
			} else {
				m[k] += string(sub)
			}
		}

		_, ok := m[string(sub)]
		if !ok {
			m[string(sub)] = string(sub)
		}
	}
	for _, v := range m {
		res = int(math.Max(float64(len(v)), float64(res)))
	}
	return res
}

// todo:1
// 16ms
func lengthOfLongestSubstring2(s string) int {
	var m = make(map[string]int)
	res := 0
	l := 0
	for r, v := range s {
		k, ok := m[string(v)]
		if ok {
			l = int(math.Max(float64(k+1), float64(l)))
		}
		res = int(math.Max(float64(res), float64(r-l+1)))
		m[string(v)] = r
	}
	return res
}
