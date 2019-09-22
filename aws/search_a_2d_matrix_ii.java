// https://leetcode.com/problems/search-a-2d-matrix-ii/ (medium)

class Solution {

    public boolean searchMatrix(int[][] matrix, int target) {
        if (matrix.length == 0  || matrix[0].length == 0) {
            return false;
        }
        for (int ci = 0; ci < matrix[0].length; ci++) {
            if (matrix[matrix.length-1][ci] >= target) {
                for (int r = matrix.length-1; r >= 0; r--) {
                    if (matrix[r][ci] == target) {
                        return true;
                    }
                }
            }
        }
        return false;
    }

}
