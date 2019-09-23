// https://leetcode.com/problems/n-ary-tree-preorder-traversal/ (easy) 

class Solution {

    public List<Integer> preorder(Node root) {
        List<Integer> l = new ArrayList<Integer>();
        if (root == null) {
            return l;
        }
        l.add(root.val);
        for (Node n : root.children) {
              l.addAll(preorder(n));
        }
        return l;
    }

}
