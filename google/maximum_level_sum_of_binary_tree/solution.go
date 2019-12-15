package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tn := &TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 18}},
	}
	fmt.Printf("expected 3, got %d\n", MaxLevelSum(tn))
}

// MaxLevelSum returns the maximum sum
// at a single depth of the tree
func MaxLevelSum(tn *TreeNode) int {
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
