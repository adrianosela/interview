// https://leetcode.com/problems/spiral-matrix-ii/ (medium)

package main

import "fmt"

func generateMatrix(n int) [][]int {
    // new n by n matrix
    // init. filled with zeros
    m := make([][]int, n)
    for i := range m {
        m[i] = make([]int, n)
    }

    p := 1
    bounds := 0

    for bounds < n {
        // write up a row
        for c := bounds; c < n-1-bounds; c++ {
            m[bounds][c] = p; p++
        }
        // write up a col
        for r := bounds; r < n-1-bounds; r++ {
            m[r][n-1-bounds] = p; p++
        }
        // write back row
        for c := n-1-bounds; c > bounds; c-- {
            m[n-1-bounds][c] = p; p++
        }
        // write back col
        for r := n-1-bounds; r > bounds; r-- {
            m[r][bounds] = p; p++
        }
        // increase bounds
        bounds++
    }

    // if odd n, write middle
    if n % 2 != 0 {
      m[n/2][n/2] = p
    }

    return m
}

func main() {
  m := generateMatrix(9)

  for _, r := range m {
    fmt.Println(r)
  }
}
