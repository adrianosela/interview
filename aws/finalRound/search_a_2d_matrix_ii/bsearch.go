package main

/*
Time complexity : O(n*log(m)) - Worst case: n binary searches
Space complexity : O(n*log(m)) - Worst case: n binary searches
*/

import "math"

// O(n*log(m)) // n binary searches
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for ri := 0; ri < len(matrix); ri++ {
		if bsearch(matrix[ri], 0, len(matrix[0])-1, target) != -1 {
			return true
		}
	}

	return false
}

// O(log(n)) - binary search in array
func bsearch(arr []int, start, end, target int) int {
	if start > end {
		return -1 // not found
	}
	mid := int(math.Ceil(float64((start + end)) / 2))
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		return bsearch(arr, start, mid-1, target)
	}
	return bsearch(arr, mid+1, end, target)
}
