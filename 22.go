package main

// String
// Backtracking

// 12ms
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	} else if n == 1 {
		return []string{"()"}
	} else {
		m := make(map[string]bool)
		for k := 1; k < n; k++ {
			before := generateParenthesis(n - k)
			if k == 1 {
				for i := 0; i < len(before); i++ {
					m["()"+before[i]] = true
					m["("+before[i]+")"] = true
					m[before[i]+"()"] = true
				}
			} else {
				after := generateParenthesis(k)
				for i := 0; i < len(before); i++ {
					for j := 0; j < len(after); j++ {
						m[before[i]+after[j]] = true
					}
				}
			}
		}

		var res []string
		for k := range m {
			res = append(res, k)
		}
		return res
	}
}

// todo:
// 8ms
func generateParenthesis2(n int) []string {
	var res []string
	return dfs(res, "", n, n)
}

func dfs(res []string, now string, l int, r int) []string {
	if l == 0 && r == 0 {
		res = append(res, now)
		return res
	}
	if l > 0 && r > 0 {
		res = dfs(res, now+"(", l-1, r)
	}
	if r > 0 && r > l {
		res = dfs(res, now+")", l, r-1)
	}
	return res
}
