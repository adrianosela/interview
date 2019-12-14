package main

/*
 Solution: Brute Force

 Runtime: O(n^2) - two 'for' loops
 Space: O(1)
*/

func maxProfit(prices []int) int {
	max := 0

	for buy := 0; buy < len(prices); buy++ {
		for sell := buy + 1; sell < len(prices); sell++ {
			profit := prices[sell] - prices[buy]
			if profit > max {
				max = profit
			}
		}
	}

	return max
}
