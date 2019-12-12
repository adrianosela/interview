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

func minDominoRotations(A []int, B []int) int {
	// we have two possible targets
	target1, target2 := A[0], B[0]

	// for each target we have a number of rotations
	// on each of A and B (top and bottom), as well as
	// an indicator for each which tells us when we
	// can no longer achieve the target through rotating
	// the top and bottom respectively
	t1TopRot, t1BottomRot, t1GiveUp := 0, 0, false
	t2TopRot, t2BottomRot, t2GiveUp := 0, 0, false

	for i := 0; i < len(A); i++ {
		top, bottom := A[i], B[i]

		// check we can achieve either target
		// by rotating either of them
		if top != target1 && bottom != target1 {
			t1GiveUp = true
		}
		if top != target2 && bottom != target2 {
			t2GiveUp = true
		}
		// if not possible on either target
		// return right away
		if t1GiveUp && t2GiveUp {
			return -1
		}

		// top rotations fir each target
		if top == target1 && bottom != target1 {
			t1BottomRot++
		}
		if top == target2 && bottom != target2 {
			t2BottomRot++
		}

		// bottom rotations for each target
		if bottom == target1 && top != target1 {
			t1TopRot++
		}
		if bottom == target2 && top != target2 {
			t2TopRot++
		}
	}

	if t1GiveUp {
		return min(t2TopRot, t2BottomRot)
	}
	return min(t1TopRot, t1BottomRot)
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
