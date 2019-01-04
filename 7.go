package main

// 4ms
func reverse(x int) int {
	var res int32 = 0
	y := int32(x)
	for y != 0 {
		k := res
		res = res*10 + y%10
		if (res-y%10)/10 != k {
			return 0
		}
		y /= 10
	}
	return int(res)
}
