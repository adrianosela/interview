// https://leetcode.com/problems/subtree-of-another-tree/ (easy)

package main

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	return areTheSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func areTheSame(s, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	return s.Val == t.Val && areTheSame(s.Left, t.Left) && areTheSame(s.Right, t.Right)
}
