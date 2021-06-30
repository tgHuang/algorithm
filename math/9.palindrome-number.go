package math

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	oldX := x
	result := 0
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
	return oldX == result
}

// @lc code=end
