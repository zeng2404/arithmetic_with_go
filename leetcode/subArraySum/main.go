package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 0, 2}
	k := 2
	result := subarraySum(nums, k)
	fmt.Println(result)
}

//func subarraySum(nums []int, k int) int {
//	top := 0
//	counter := 0
//	for top < len(nums) {
//		tail := top
//		count := 0
//		for tail < len(nums) {
//			if count += nums[tail]; count == k {
//				counter++
//			}
//			tail++
//		}
//		top++
//	}
//	return counter
//}

//func subarraySum(nums []int, k int) int {
//	count := 0
//	for top:= 0; top < len(nums); top++ {
//		sum := 0
//		for tail := top; tail >= 0; tail-- {
//			sum += nums[tail]
//			if sum == k {
//				count++
//			}
//		}
//	}
//	return count
//}

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := make(map[int]int)
	// 假如第一个 pre 就是符合要求的子元素，那么就要设置要 1，防止第一个次命中丢失一次计数
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}
