// https://leetcode.com/problems/subtree-of-another-tree/ (easy)

package main

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}
	return areTheSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func areTheSame(s, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}
	if t == nil {
		return false
	}
	return s.Val == t.Val && areTheSame(s.Left, t.Left) && areTheSame(s.Right, t.Right)
}
