package main

import "fmt"

func main() {
	pushed := []int{0, 2, 1}
	popped := []int{0, 1, 2}

	fmt.Println(validateStackSequences(pushed, popped))

}

func validateStackSequences(pushed []int, popped []int) bool {
	// 创建一个辅助栈
	stack := make([]int, len(pushed), len(pushed))

	top := 0
	j := 0
	for i := 0; i < len(pushed); i++ {
		// 入栈
		stack[top] = pushed[i]
		top++
		// 当满足出栈条件，则出栈
		for top > 0 && stack[top-1] == popped[j] {
			top--
			j++
		}

	}
	// 当元素全部入栈，但是却未能全部出栈时，则表示出栈结果错误
	return top == 0

}
