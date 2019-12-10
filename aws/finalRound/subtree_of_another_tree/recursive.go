// https://leetcode.com/problems/subtree-of-another-tree/ (easy)

/*
Time complexity : O(m*n) - In worst case (skewed tree)
Space complexity : O(n) - Recursion depth can go up to n where n = nodes in s
*/

package main

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s, t *TreeNode) bool {
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
