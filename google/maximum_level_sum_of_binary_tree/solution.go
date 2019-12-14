package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxLevelSum returns the maximum sum
// at a single depth of the tree
func maxLevelSum(tn *TreeNode) int {
	// need a queue to do bfs / level-order
	q := []*TreeNode{tn}

	// we must keep track of the
	// maximum sum 'so far'
	max, ans := math.MinInt32, 1

	var next *TreeNode
	for lvl := 1; len(q) > 0; lvl++ {
		sum := 0

		// note that i's value is the
		// size of the queue BEFORE
		// the loop adds any items
		for i := len(q); i > 0; i-- {
			next, q = q[0], q[1:]
			sum += next.Val

			if next.Left != nil {
				q = append(q, next.Left)
			}
			if next.Right != nil {
				q = append(q, next.Right)
			}
		}

		if sum > max {
			ans = lvl
			max = sum
		}
	}

	return ans

}
