package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, -2))
}

//func myPow(x float64, n int) float64 {
//	isMinis := n < 0
//	var sum float64 = 1
//	if isMinis {
//		n = -n
//	}
//	for i := 0; i < n; i++ {
//		sum *= x
//	}
//	if isMinis {
//		return 1 / sum
//	} else {
//		return sum
//	}
//}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var sum float64 = 1
	count := 0
	// 表示 n 是否为 负数
	isMinus := n < 0
	if isMinus {
		n = -n
	}

	// 用于记录平方次数的值
	arr := make([]float64, 1, 1)
	arr[0] = x

	// 将平方倍记录下来，转换为位运算
	for i := 2; i <= n; i = i * 2 {
		arr = append(arr, arr[len(arr)-1]*arr[len(arr)-1])
	}

	for n != 0 {
		if n&1 == 1 {
			sum *= arr[count]
		}
		n = n >> 1
		count++
	}

	if isMinus {
		return 1 / sum
	} else {
		return sum
	}
}
