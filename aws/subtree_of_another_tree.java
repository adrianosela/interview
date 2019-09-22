// https://leetcode.com/problems/subtree-of-another-tree/ (easy)

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
