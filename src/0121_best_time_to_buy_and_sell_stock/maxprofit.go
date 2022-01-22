/*
121. Best Time to Buy and Sell Stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction
(i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.
*/
// time: 2018-12-28

package maxprofit

// dynamic programming
// for every price, find the max profit, then record the current minimum price.
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

	for _, price := range prices {
		if price-minPrice > res {
			res = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return res
}

// brute force
// time complexity: O(n^2)
// space complexity: O(1)
func maxProfit1(prices []int) int {
	n := len(prices)
	if 0 == n || 1 == n {
		return 0
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > res {
				res = prices[i] - prices[j]
			}
		}
	}
	return res
}
