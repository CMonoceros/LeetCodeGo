package main

import "strconv"

// String
// Backtracking

func letterMap(num int) []string {
	switch num {
	case 2:
		return []string{"a", "b", "c"}
	case 3:
		return []string{"d", "e", "f"}
	case 4:
		return []string{"g", "h", "i"}
	case 5:
		return []string{"j", "k", "l"}
	case 6:
		return []string{"m", "n", "o"}
	case 7:
		return []string{"p", "q", "r", "s"}
	case 8:
		return []string{"t", "u", "v"}
	case 9:
		return []string{"w", "x", "y", "z"}
	}
	return []string{}
}

// 0ms
func letterCombinations(digits string) []string {
	d := []rune(digits)

	if len(d) == 0 {
		return []string{}
	} else if len(d) == 1 {
		num, _ := strconv.Atoi(string(d[:1]))
		return letterMap(num)
	} else {
		num, _ := strconv.Atoi(string(d[:1]))
		before := letterMap(num)
		after := letterCombinations(string(d[1:]))
		var res []string

		for i := 0; i < len(before); i++ {
			for j := 0; j < len(after); j++ {
				res = append(res, before[i]+after[j])
			}
		}
		return res
	}
}

// todo:1
// 0ms
func letterCombinations2(digits string) []string {
	d := []rune(digits)
	if len(d) == 0 {
		return []string{}
	}
	var dfs func(res []string, now string, d []rune) []string
	dfs = func(res []string, now string, d []rune) []string {
		if len(now) == len(d) {
			res := append(res, now)
			return res
		}
		num, _ := strconv.Atoi(string(d[len(now)]))
		m := letterMap(num)
		for i := 0; i < len(m); i++ {
			res = dfs(res, now+m[i], d)
		}
		return res
	}
	return dfs([]string{}, "", d)
}
