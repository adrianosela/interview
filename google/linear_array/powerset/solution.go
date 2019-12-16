// https://leetcode.com/problems/subsets/submissions/

package main

func subsets(nums []int) [][]int {
	output := [][]int{{}}
	for _, elem := range nums {
		for i := len(output) - 1; i >= 0; i-- {
			tmp := make([]int, len(output[i]))
			copy(tmp, output[i])
			tmp = append(tmp, elem)
			output = append(output, tmp)
		}
	}
	return output
}
