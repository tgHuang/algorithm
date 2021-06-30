package math

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100,
		"D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}

	var result, lenth = 0, len(s)
	for i := 0; i < lenth; i++ {
		if i == lenth-1 {
			result += m[string(s[i])]
			continue
		}
		two := s[i : i+2]
		switch two {
		case "IV", "IX", "XL", "XC", "CD", "CM":
			result += m[two]
			i++
			continue
		default:
			result += m[string(s[i])]
		}
	}
	return result
}

// @lc code=end
