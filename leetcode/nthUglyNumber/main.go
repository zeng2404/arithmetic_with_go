package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("the value is %v\n", nthUglyNumber(11))
}

func min(a int, b int, c int) int {
	return int(math.Min(math.Min(float64(a), float64(b)), float64(c)))
}

func nthUglyNumber(n int) int {
	arr := make([]int, 1, 1)
	arr[0] = 1

	a, b, c := 0, 0, 0

	for len(arr) < n {
		// 去重
		if arr[a]*2 == arr[len(arr)-1] {
			a++
		}
		if arr[b]*3 == arr[len(arr)-1] {
			b++
		}
		if arr[c]*5 == arr[len(arr)-1] {
			c++
		}

		// 比较最小值
		minValue := min(arr[a]*2, arr[b]*3, arr[c]*5)
		arr = append(arr, minValue)
	}
	return arr[n-1]
}
