/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public int maxLevelSum(TreeNode root) {
        
        // need a queue to do BFS / level order traversal
        Queue<TreeNode> q = new LinkedList<TreeNode>();
        q.add(root);
        
        // need to keep track of max 
        // sum so far and the result
        int max = Integer.MIN_VALUE;
        int ans = 1;
        
       for (int level = 1; !q.isEmpty(); level++) {
            int size = q.size();
            int sum = 0;
            
            for (int i = 0; i < size; i++) {
                TreeNode next = q.remove();
                sum += next.val;
                if (next.left != null) { q.add(next.left); }
                if (next.right != null) { q.add(next.right); }
            }
            
            if (sum > max) {
                ans = level;
                max = sum;
            }
        }
        
        return ans;
    }
}
