package main

import (
	"fmt"
	"math"
)

func main() {
	nums := make([]int, 3)
	nums = []int{-2}
	fmt.Printf("the max product is %d\n", maxProduct(nums))
}

//func maxProduct(nums []int) int {
//	maxValue := nums[0]
//
//	for i := 0; i < len(nums); i++ {
//		currentValue := 1
//		for j := i; j < len(nums); j++ {
//			currentValue *= nums[j]
//			if currentValue > maxValue {
//				maxValue = currentValue
//			}
//		}
//	}
//
//	return maxValue
//}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxProduct(nums []int) int {
	maxValue := 1
	minValue := 1
	currentMax := -int(^uint(0)>>1) - 1

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			temp := minValue
			minValue = maxValue
			maxValue = temp
		}

		maxValue = max(maxValue*nums[i], nums[i])
		minValue = min(minValue*nums[i], nums[i])

		currentMax = max(currentMax, maxValue)
	}

	return currentMax
}
