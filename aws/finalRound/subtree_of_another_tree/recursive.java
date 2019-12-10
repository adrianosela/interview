// https://leetcode.com/problems/subtree-of-another-tree/ (easy)

/*
Time complexity : O(m*n) - In worst case (skewed tree)
Space complexity : O(n) - Recursion depth can go up to n where n = nodes in s
*/

class Solution {
    public boolean isSubtree(TreeNode s, TreeNode t) {
        if(s == null) {
            return false;
        }
        if (s.val != t.val) {
            return isSubtree(s.left, t) || isSubtree(s.right, t);
        }
        return check(s, t)
                || isSubtree(s.left, t)
                || isSubtree(s.right, t);
    }
    private boolean check(TreeNode s, TreeNode t) {
        if(s == null) {
            return t == null;
        } else {
            if (t == null) {
                return false;
            }
        }
        return s.val == t.val && check(s.left, t.left) && check(s.right, t.right);
    }
}
