package array

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		if v, ok := m[target-val]; ok {
			return []int{v, i}
		}
		m[val] = i
	}
	return []int{}
}

// @lc code=end
