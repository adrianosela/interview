// https://leetcode.com/problems/search-a-2d-matrix-ii/ (medium)
package main

/*
Time complexity : O(m+n) - Worst case: full pass in each dimension (rows & columns)
Space complexity : O(m+n) - Worst case: full pass in each dimension (rows & columns)
*/

// search matrix starting from the bottom left
func searchMatrixBL(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	i := len(matrix) - 1
	j := 0

	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
			continue
		}
		if matrix[i][j] < target {
			j++
			continue
		}
		return true
	}

	return false
}

// search matrix starting from the top right
func searchMatrixTR(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	i := 0
	j := len(matrix[0]) - 1

	for i < len(matrix) && j >= 0 {
		if matrix[i][j] < target {
			i++
			continue
		}
		if matrix[i][j] > target {
			j--
			continue
		}
		return true
	}

	return false
}
