package main

// from bottom left
func searchMatrix(matrix [][]int, target int) bool {
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
