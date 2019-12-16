// https://leetcode.com/problems/subarray-sum-equals-k

package main

/*
Time complexity : O(n) - All cases: one full pass in the array
Space complexity : O(n) - All cases: we keep an n-sized map
*/

func subarraySum(nums []int, k int) int {
	sums, cur, upTo := 0, 0, map[int]int{0: 1}

	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		if prevSums, ok := upTo[cur-k]; ok {
			sums += prevSums
		}

		if _, ok := upTo[cur]; ok {
			upTo[cur]++
		} else {
			upTo[cur] = 1
		}
	}

	return sums
}
