package dynamic

// 121. Best Time to Buy and Sell Stock
// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
//
//
// Example 1:
//
// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:
//
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.
//
//
// Constraints:
//
// 1 <= prices.length <= 1e5
// 0 <= prices[i] <= 1e4
// 股票利润最大化售卖时间
func maxProfit(prices []int) int {
	var min int = 1e4
	var maxProfit int = 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if v-min > maxProfit {
			maxProfit = v - min
		}
	}
	//	lenth := len(prices)
	//	for i := 0; i < lenth; i++ {
	//		if prices[i] < min {
	//			min = prices[i]
	//		} else if prices[i]-min > maxProfit {
	//			maxProfit = prices[i] - min
	//		}
	//	}
	return maxProfit
}

// 122. Best Time to Buy and Sell Stock II
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
//
//
// Example 1:
//
// Input: prices = [7,1,5,3,6,4]
// Output: 7
// Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
// Example 2:
//
// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
// Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.
// Example 3:
//
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e., max profit = 0.
//
//
// Constraints:
//
// 1 <= prices.length <= 3 * 1e4
// 0 <= prices[i] <= 1e4
func maxProfit2(prices []int) int {
	//	var lastVal, profit int
	//	for i, v := range prices {
	//		if i == 0 {
	//			lastVal = v
	//			continue
	//		}
	//		if v >= lastVal {
	//			profit = profit + v - lastVal
	//		}
	//		lastVal = v
	//	}

	if len(prices) <= 0 {
		return 0
	}
	var lastVal, profit, lenth int = prices[0], 0, len(prices)

	for i := 1; i < lenth; i++ {
		if prices[i] >= lastVal {
			profit = profit + prices[i] - lastVal
		}
		lastVal = prices[i]
	}

	return profit
}

// 123. Best Time to Buy and Sell Stock III
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// Find the maximum profit you can achieve. You may complete at most two transactions.
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
//
//
// Example 1:
//
// Input: prices = [3,3,5,0,0,3,1,4]
// Output: 6
// Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
// Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
// Example 2:
//
// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
// Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.
// Example 3:
//
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e. max profit = 0.
// Example 4:
//
// Input: prices = [1]
// Output: 0
//
//
// Constraints:
//
// 1 <= prices.length <= 1e5
// 0 <= prices[i] <= 1e5
// 代码是否可以抽象出来？
func maxProfit3(prices []int) int {
	lenth := len(prices)
	if lenth <= 1 {
		return 0
	}
	// lastMin mean except two before point
	var lastVal, min1, max1, min2, max2 int = prices[0], prices[0], 0, 0, 0
	var pro1, pro2 int = 0, 0
	var lastMin int = -1
	for i := 1; i < lenth; i++ {
		// down
		if prices[i] < lastVal {
			// first did not buy
			if pro1 == 0 {
				lastVal = prices[i]
				continue
			}
			// first buy, but second did not buy
			if pro2 == 0 {
				if lastVal > max1 {
					max1 = lastVal
					pro1 = max1 - min1
				}
				if min2 == 0 || min2 > prices[i] {
					min2 = prices[i]
				}
				lastVal = prices[i]
				continue
			}
			// first and second all buy, record third pro
			if lastMin == -1 {
				lastMin = prices[i]
				lastVal = prices[i]
				continue
			}

			if lastVal-min2 > pro2 && lastVal-min2 > lastVal-lastMin {
				max2 = lastVal
				pro2 = max2 - min2
				lastMin = -1
				lastVal = prices[i]
				continue
			}

			if pro1 >= pro2 && lastVal-lastMin > pro2 {
				if max2 > max1 {
					max1 = max2
					pro1 = max2 - min2
				}
				max2 = lastVal
				min2 = lastMin
				pro2 = max2 - min2
				lastMin = -1
				lastVal = prices[i]
				continue
			}

			if pro1 < pro2 && lastVal-lastMin > pro1 {
				if min1 >= min2 {
					min1 = min2
				}
				max1 = max2
				pro1 = max1 - min1
				min2 = lastMin
				max2 = lastVal
				pro2 = max2 - min2
				lastMin = -1
				lastVal = prices[i]
				continue
			}
			if lastMin > prices[i] {
				lastMin = prices[i]
			}
		}
		// up
		if prices[i] >= lastVal {
			if pro1 == 0 {
				min1 = lastVal
				max1 = prices[i]
				pro1 = max1 - min1
				lastVal = prices[i]
				continue
			}

			if lastVal == max1 && pro2 == 0 {
				max1 = prices[i]
				pro1 = max1 - min1
				lastVal = prices[i]
				continue
			}

			if pro2 == 0 {
				min2 = lastVal
				max2 = prices[i]
				pro2 = max2 - min2
				lastVal = prices[i]
				continue
			}
			if prices[i] > max2 && lastVal == max2 {
				max2 = prices[i]
				pro2 = max2 - min2
				lastVal = prices[i]
				continue
			}
			if lastMin == -1 {
				lastMin = lastVal
				lastVal = prices[i]
				continue
			}
			// 小心：这儿会吞掉更大的值
			if prices[i]-min2 > pro2 && prices[i]-min2+pro1 >= prices[i]-lastMin+pro2 {
				if min2 > lastMin {
					min2 = lastMin
				}
				max2 = prices[i]
				pro2 = max2 - min2
				lastMin = -1
				lastVal = prices[i]
				continue
			}

			if pro1 >= pro2 && prices[i]-lastMin > pro2 {
				if max2 > max1 {
					max1 = max2
					pro1 = max2 - min1
				}
				max2 = prices[i]
				min2 = lastMin
				lastMin = -1
				pro2 = max2 - min2
				lastVal = prices[i]
				continue
			}

			if pro1 < pro2 && prices[i]-lastMin > pro1 {
				if min1 >= min2 {
					min1 = min2
				}
				max1 = max2
				pro1 = max1 - min1
				min2 = lastMin
				lastMin = -1
				max2 = prices[i]
				pro2 = max2 - min2
				lastVal = prices[i]
				continue
			}

		}

		lastVal = prices[i]
	}
	return pro1 + pro2
}

// 188. Best Time to Buy and Sell Stock IV
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// You are given an integer array prices where prices[i] is the price of a given stock on the ith day, and an integer k.
//
// Find the maximum profit you can achieve. You may complete at most k transactions.
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
// Example 1:
//
// Input: k = 2, prices = [2,4,1]
// Output: 2
// Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
// Example 2:
//
// Input: k = 2, prices = [3,2,6,5,0,3]
// Output: 7
// Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4. Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
//
// Constraints:
//
// 0 <= k <= 100
// 0 <= prices.length <= 1000
// 0 <= prices[i] <= 1000
func maxProfit4(k int, prices []int) int {
	var lastVal, upMax, upMin, curOp int = prices[0], -1, -1, 0
	var lastUpMin, lastUpMax int = -1, -1
	ops := make([]int, 0)
	for i := 1; i < len(prices); i++ {
		// up
		if prices[i] > lastVal {

			if lastVal == upMax {
				upMax = prices[i]
				lastVal = prices[i]
				continue
			}

			if upMin == -1 || upMax == -1 {
				upMin = lastVal
				upMax = prices[i]
				continue
			}

			curOp = upMax - upMin
			if len(ops) < k {
				lastUpMin = upMin
				lastUpMax = upMax
				upMin = lastVal
				upMax = prices[i]
				ops = append(ops, curOp)
				continue
			}
			upMin = lastVal
			upMax = prices[i]
			// 达到操作上限
			sum, min := sumWithOutMin(ops)
			if upMax-upMin < min {
				if upMax > lastUpMax {
					lastUpMax = upMax
					continue
				}
			}
		}
		lastVal = prices[i]
	}
	curOp = upMax - upMin
	if curOp > 0 && len(ops) < k {
		ops = append(ops, curOp)
	}

	var total int
	for _, v := range ops {
		total += v
	}
	return total
}

func sumWithOutMin(ops []int) (sum, min int) {
	if len(ops) == 0 {
		return
	}
	min, sum = ops[0], 0
	for i := 1; i < len(ops); i++ {
		if ops[i] < min {
			sum += min
			min = ops[i]
			continue
		}
		sum += ops[i]
	}
	return
}
