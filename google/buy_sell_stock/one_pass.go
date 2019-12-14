package main

import "math"

/*
 Solution: One Pass, we keep a min and a max
 variables and return the largest difference
 at the current index

 Runtime: O(n) - one 'for' loop
 Space: O(1)
*/

func maxProfit(prices []int) int {
	min := math.MaxInt32
	max := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > max {
				max = prices[i] - min
			}
		}
	}

	return max
}
