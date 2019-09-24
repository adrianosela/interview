// https://leetcode.com/problems/evaluate-reverse-polish-notation/ (medium)

class Solution {

    public int evalRPN(String[] tokens) {
        Stack<Integer> s = new Stack();
        for (int i = 0; i < tokens.length; i++) {
            String t = tokens[i];
            int a, b;
            switch (tokens[i]) {
                case "+":
                    s.push(s.pop()+s.pop());
                    break;
                case "-":
                    b = s.pop();
                    a = s.pop();
                    s.push(a-b);
                    break;
                case "*":
                    s.push(s.pop()*s.pop());
                    break;
                case "/":
                    b = s.pop();
                    a = s.pop();
                    s.push(a/b);
                    break;
                default:
                    s.push(Integer.parseInt(t));
            }
        }
        return s.pop(); // only result prevails
    }

}
