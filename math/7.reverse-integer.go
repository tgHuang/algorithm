package math

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	max, min := 2<<30-1, -2<<30
	result := 0
	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}

	for x > 0 {
		rem := x % 10
		x = x / 10
		if result == 0 {
			if rem > 0 {
				result = rem
			}
			continue
		}
		result = result*10 + rem
	}

	if isNegative {
		result = -result
		if result < min {
			return 0
		}
		return result
	}
	if result > max {
		return 0
	}
	return result
}

// @lc code=end
