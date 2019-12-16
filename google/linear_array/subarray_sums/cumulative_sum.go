// https://leetcode.com/problems/subarray-sum-equals-k

package main

/*
Time complexity : O(n^2) - All cases: two full for-loops, no breaking
Space complexity : O(1) - We look at one item at once
*/

func subarraySum(nums []int, k int) int {
	sums := 0
	for i := 0; i < len(nums); i++ {
		cur := 0
		for j := i; j < len(nums); j++ {
			cur += nums[j]
			if cur == k {
				sums++
			}
		}
	}

	return sums
}
