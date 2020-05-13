package main

import "fmt"

func main() {
	var board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	e := exist(board, "ABCESEEEFS")
	fmt.Println(e)
}

// 原来特别慢的解法
//func exist(board [][]byte, word string) bool {
//	if len(word) == 0 {
//		return true
//	}
//	// 获取长和宽
//	length := len(board[0])
//	width := len(board)
//
//	// 维护一个已经路过的节点的数组
//	a := make([][]int, width)
//	for k := 0; k < width; k++ {
//		array := make([]int, length)
//		a[k] = array
//	}
//
//	// 维护一个任务队列
//	//jobs := make([]element, 0)
//
//	flag := false
//
//	var f func(index int, x int, y int)
//
//	f = func(index int, x int, y int) {
//		if index >= len(word) || flag == true {
//			flag = true
//			return
//		}
//		// 节点标记为已经路过
//		a[y][x] = 1
//		// 判断节点的四个方向是否已经路过或者超出边界
//		if y-1 >= 0 && a[y-1][x] != 1 {
//			// 匹配成果深度加一
//			if board[y-1][x] == word[index] {
//				f(index+1, x, y-1)
//			}
//		}
//
//		if y+1 < width && a[y+1][x] != 1 {
//			if board[y+1][x] == word[index] {
//				f(index+1, x, y+1)
//			}
//		}
//
//		if x-1 >= 0 && a[y][x-1] != 1 {
//			if board[y][x-1] == word[index] {
//				f(index+1, x-1, y)
//			}
//		}
//
//		if x+1 < length && a[y][x+1] != 1 {
//			if board[y][x+1] == word[index] {
//				f(index+1, x+1, y)
//			}
//		}
//		a[y][x] = 0
//	}
//
//	for i := 0; i < width; i++ {
//		for j := 0; j < length; j++ {
//			if board[i][j] == word[0] {
//				a = make([][]int, width)
//				for k := 0; k < width; k++ {
//					array := make([]int, length)
//					a[k] = array
//				}
//				f(1, j, i)
//			}
//			if flag {
//				return true
//			}
//		}
//	}
//	return false
//}

// 修改后的解法
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i int, j int, k int) bool {
	// 不需要直接找到第一个点，而是直接针对一个点进行判断，
	//如果当前的 k 值没有达到字符串的长度且和现在的层级所在字符不匹配，则直接返回 false
	// 以及下标超界了也直接返回 false
	// 如果 k 值达到了 字符串的长度，则返回 true
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	if k == (len(word) - 1) {
		return true
	}
	temp := board[i][j]
	// 替换成字符无法匹配的符号，也就是做了标记
	board[i][j] = '/'
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i, j+1, k+1) || dfs(board, word, i, j-1, k+1)

	// 在回溯的时候需要把原来标记走过的元素恢复
	board[i][j] = temp
	return res
}
