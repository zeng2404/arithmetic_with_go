package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(cuttingRope(10))
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

//func cuttingRope(n int) int {
//	if n == 2 {
//		return 1
//	}
//	res := -1
//	for i := 1; i < n; i++ {
//		res = max(res, max(i * cuttingRope(n - i), i * (n - i)))
//	}
//	return res
//}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	dp := make([]int, n+1, n+1)
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		if 2*dp[i-2] > 3*dp[i-3] {
			dp[i] = 2 * dp[i-2]
		} else {
			dp[i] = 3 * dp[i-3]
		}
	}
	return dp[n]
}
