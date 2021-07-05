package array

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// odd len / 2 is midian, (even len / 2) + (even len / 2 - 1) / 2 is midian
	m, n := len(nums1), len(nums2)
	isOdd := (m+n)%2 != 0
	if m == 0 {
		if n%2 == 0 {
			return (float64(nums2[n/2]) + float64(nums2[n/2-1])) / 2
		}
		return float64(nums2[n/2])
	}
	if n == 0 {
		if m%2 == 0 {
			return (float64(nums1[m/2]) + float64(nums1[m/2-1])) / 2
		}
		return float64(nums1[m/2])
	}
	var pos1, pos2, max1, min1, max2, min2 = 0, 0, m, 0, n, 0
	// start midian,start less one
	for {

		if isOdd && pos1+pos2 >= (m+n)/2 {
			break
		}
		if nums1[pos1] == nums2[pos2] {
			pos1 += 1
			pos2 += 1
			min1, min2 = pos1, pos2
			continue
		}
		if nums1[pos1] > nums2[pos2] {
			pos2 = (max2 - min2) / 2
			min2 = pos2
		}
		if nums1[pos1] < nums2[pos2] {
			pos1 = (max1 - min1) / 2
			min1 = pos1
		}
	}
}

// @lc code=end
