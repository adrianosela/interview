package main

/*
--------------------- PROBLEM: --------------------------------
- In a row of dominoes, A[i] and B[i] represent the top and
  bottom halves of the i-th domino.  (A domino is a tile with
  two numbers from 1 to 6 - one on each half of the tile.)

- We may rotate the i-th domino, so that A[i] and B[i]
  swap values.

- Return the minimum number of rotations so that all the
  values in A are the same, or all the values in B are the same.

- If it cannot be done, return -1.
--------------------------------------------------------------
*/

import "fmt"

func main() {
	run1 := minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2})
	fmt.Printf("minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}) - expected 2, got %d\n", run1)

	run2 := minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4})
	fmt.Printf("minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}) - expected -1, got %d\n", run2)
}

func minDominoRotations(A, B []int) int {
	if numA := swapOnTarget(A[0], A, B); numA != -1 {
		return numA
	}
	if numB := swapOnTarget(B[0], B, A); numB != -1 {
		return numB
	}
	return -1
}

func swapOnTarget(target int, A, B []int) int {
	a, b := 0, 0
	for i := 0; i < len(A); i++ {
		if A[i] != target && B[i] != target {
			return -1
		}

		if B[i] == target && A[i] != target {
			a++
		}

		if A[i] == target && B[i] != target {
			b++
		}
	}
	return min(a, b)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
