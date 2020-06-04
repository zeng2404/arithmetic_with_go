package main

import "fmt"

// https://leetcode-cn.com/problems/product-of-array-except-self/

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(arr))
}

func productExceptSelf(nums []int) []int {
	// 创建左数组
	l := make([]int, len(nums)-1, len(nums)-1)
	r := make([]int, len(nums)-1, len(nums)-1)
	l[0] = nums[0]
	r[len(nums)-2] = nums[len(nums)-1]

	for i := 1; i < len(nums)-1; i++ {
		l[i] = l[i-1] * nums[i]
	}

	for j := len(nums) - 3; j >= 0; j-- {
		r[j] = r[j+1] * nums[j+1]
	}

	// 要返回的数组
	result := make([]int, len(nums), len(nums))
	result[0] = r[0]
	result[len(result)-1] = l[len(result)-2]
	for k := 1; k <= len(result)-2; k++ {
		result[k] = r[k] * l[k-1]
	}

	return result
}
