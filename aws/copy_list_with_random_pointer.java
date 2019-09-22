// https://leetcode.com/problems/copy-list-with-random-pointer/ (medium)

class Solution { 
    
    private Map<Node, Node> copies = new HashMap<Node,Node>();
    
    public Node copyRandomList(Node head) {
        if (head == null) {
            return null;
        }
        
        Node cpy = new Node(head.val, null, null);
        copies.put(head, cpy);
        
        if (head.next != null) {
            cpy.next = copyRandomList(head.next);
        }
        
        // at this point all nodes are in map
        if (head.random != null) {
            cpy.random = copies.get(head.random);
        }
        
        return cpy;
    }
    
}
