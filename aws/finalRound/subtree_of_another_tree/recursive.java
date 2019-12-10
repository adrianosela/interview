// https://leetcode.com/problems/subtree-of-another-tree/ (easy)

/*
Time complexity : O(m*n) - In worst case (skewed tree)
Space complexity : O(n) - Recursion depth can go up to n where n = nodes in s
*/

class Solution {
    public boolean isSubtree(TreeNode s, TreeNode t) {
        if (s == null || t == null) {
            return s == t;
        }
        return areTheSame(s, t)
                || isSubtree(s.left, t)
                || isSubtree(s.right, t);
    }

    private boolean areTheSame(TreeNode s, TreeNode t) {
        if (s == null || t == null) {
            return s == t;
        }
        return s.val == t.val
                && areTheSame(s.left, t.left)
                && areTheSame(s.right, t.right);
    }
}
