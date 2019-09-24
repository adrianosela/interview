// https://leetcode.com/problems/k-th-symbol-in-grammar/ (medium)

package main

func main() {
  kthGrammar(5, 8)
}

func kthGrammar(N int, K int) int {
    if N == 0 {
        return 0
    }
    if K % 2 == 1 {
        return kthGrammar(N-1, (K+1)/2)
    }
    return 1^kthGrammar(N-1, K/2)
}


// // still uses too much memory
// func kthGrammar(N int, K int) int {
//     if N <= 1 {
//         return 0
//     }
//     reverseHalf := func(a []int) {
//       for i := len(a)/2; i < len(a); i++ {
//         if a[i] == 0 {
//           a[i] = 1
//         } else {
//           a[i] = 0
//         }
//       }
//     }
//     fw := []int{0,1}
//     for i := 2; i <= N; i++ {
//       fw = append(fw, fw...)
//       reverseHalf(fw)
//       fmt.Printf("i=%d, %v\n", i, fw)
//     }
//     return fw[K-1]
// }

// // below works, but uses too much memory
// func kthGrammar(N int, K int) int {
//     if N <= 1 {
//         return 0
//     }
//     fw := []int{0,1}
//     bw := []int{1,0}
//     for i := 2; i <= N; i++ {
//       fw, bw = append(fw, bw...), append(bw, fw...)
//       fmt.Printf("i=%d, %v\n", i, fw)
//     }
//     return fw[K-1]
// }
