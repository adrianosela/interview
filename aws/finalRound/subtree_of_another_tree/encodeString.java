// https://leetcode.com/problems/subtree-of-another-tree/ (easy)

/*
Time complexity : O(m^2+n^2+m*n) - A total of n nodes of the tree s and m nodes of tree t are traversed.
		Assuming string concatenation takes O(k) time for strings of length k and indexOf takes O(m*n)
Space complexity : O(max(m,n)) - The depth of the recursion tree can go upto n for tree t and m for tree s in worst case. 
*/

class Solution {
    public boolean isSubtree(TreeNode s, TreeNode t) {
        String cmpTree = encodeTree(s);
        String subTree = encodeTree(t);
        return cmpTree.indexOf(subTree) > -1;
    } 

    // encodeTree encodes a tree into a nested string
    // representation of its preorder traversal
    public String encodeTree(TreeNode t) {
        if (t == null) {
                return "";
        }
        return "val(" + t.val + ")"
            + "left(" + encodeTree(t.left) + ")"
            + "right(" + encodeTree(t.right) + ")";
    }
}
