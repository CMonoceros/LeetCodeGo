package main

// Array
// Binary Search
// Divide and Conquer

// 28 ms
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res []int
	i, j := 0, 0
	for {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] > nums2[j] {
				res = append(res, nums2[j])
				j++
			} else {
				res = append(res, nums1[i])
				i++
			}
		} else if i < len(nums1) && j >= len(nums2) {
			res = append(res, nums1[i])
			i++
		} else if i >= len(nums1) && j < len(nums2) {
			res = append(res, nums2[j])
			j++
		} else {
			break
		}
	}
	l := len(res)
	if l%2 == 0 {
		return float64(float64(res[l/2-1]+res[l/2]) / 2)
	} else {
		return float64(res[l/2])
	}
}
