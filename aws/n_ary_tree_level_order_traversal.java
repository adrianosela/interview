// https://leetcode.com/problems/n-ary-tree-level-order-traversal/ (easy)

class Solution {

    public List<List<Integer>> levelOrder(Node root) {
        List<List<Integer>> soln = new ArrayList<List<Integer>>();
        if (root == null){
            return soln;
        }

        Queue<Node> queue = new LinkedList<Node>();
        queue.add(root);

        while (!queue.isEmpty()) {
            int size = queue.size();
            List<Integer> level = new ArrayList();
            for (int i = 0; i < size; ++i) {
                Node node = queue.remove();
                level.add(node.val);
                for (Node child : node.children) {
                    queue.add(child);
                }
            }
            soln.add(level);
        }
        return soln;
    }

}
