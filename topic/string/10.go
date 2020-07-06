package main

// String
// Dynamic Programming
// Backtracking

// todo:
// 288ms
func isMatch(s1 string, s2 string) bool {
	s := []rune(s1)
	p := []rune(s2)
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) != 0 && (s[0] == p[0] || string(p[0]) == ".")
	if len(p) > 1 && string(p[1]) == "*" {
		return isMatch(string(s), string(p[2:])) ||
			firstMatch && isMatch(string(s[1:]), string(p))
	} else {
		return firstMatch && isMatch(string(s[1:]), string(p[1:]))
	}
}

// 通过dp数组缓存递归结果
// todo:
// 4ms
func isMatch2(s1 string, s2 string) bool {
	s := []rune(s1)
	p := []rune(s2)
	var dp = make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := i < len(s) && (s[i] == p[j] || string(p[j]) == ".")
			if j+1 < len(p) && string(p[j+1]) == "*" {
				dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}
