// https://leetcode.com/problems/search-a-2d-matrix-ii/ (medium)

/*
Time complexity : O(m*n) - Worst case: visit every entry in the matrix
Space complexity : O(m*n) - Worst case: visit every entry in the matrix
*/

class Solution {

    public boolean searchMatrix(int[][] matrix, int target) {
        if (matrix.length == 0  || matrix[0].length == 0) {
            return false;
        }
        for (int ci = 0; ci < matrix[0].length; ci++) {
            for (int r = matrix.length-1; r >= 0; r--) {
                if (matrix[r][ci] == target) {
                    return true;
                }
            }
        }
        return false;
    }

}
