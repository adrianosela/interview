package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxLevelSum returns the maximum sum
// at a single depth of the tree
func MaxLevelSum(tn *TreeNode) int {
	// need to store depth along the
	// tree nodes somehow, so create
	// a new type
	type lTreeNode struct {
		n *TreeNode
		d int
	}

	// need queue for level order
	// traversal / BFS of tree
	q := []*lTreeNode{{n: tn, d: 0}}
	var next *lTreeNode

	// will need to keep track of
	// the current maximum sum, the current
	// level sum, and the current depth
	max, cur, depth := 0, 0, 0

	for len(q) > 0 {
		next, q = q[0], q[1:] // pop one item off the queue

		if next.d == depth {
			cur += next.n.Val
		} else {
			// update max if greater
			if cur > max {
				max = cur
			}
			// update level
			cur = next.n.Val
			depth = next.d
		}

		if next.n.Left != nil {
			q = append(q, &lTreeNode{n: next.n.Left, d: next.d + 1})
		}
		if next.n.Right != nil {
			q = append(q, &lTreeNode{n: next.n.Right, d: next.d + 1})
		}

	}

	return max
}

func main() {
	tn := &TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7}},
		Right: &TreeNode{Val: 12, Left: &TreeNode{Val: 11}, Right: &TreeNode{Val: 13}},
	}
	fmt.Println(MaxLevelSum(tn))
}
