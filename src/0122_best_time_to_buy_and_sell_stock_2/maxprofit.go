/*
122. Best Time to Buy and Sell Stock II
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit.
You may complete as many transactions as you like
(i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time
(i.e., you must sell the stock before you buy again).
*/
// time: 2018-12-28

package maxprofit

// greedy
// time complexity: O(n)
// space complexity: O(1)
func maxProfit(prices []int) int {
	n := len(prices)

	if 0 == n || 1 == n {
		return 0
	}

	var (
		res      int
		minPrice = prices[0]
	)
	for i := 1; i < n; i++ {
		if prices[i] < prices[i-1] {
			res += prices[i-1] - minPrice
			minPrice = prices[i]
		}
		if i == n-1 {
			res += prices[i] - minPrice
		}
	}
	return res
}
