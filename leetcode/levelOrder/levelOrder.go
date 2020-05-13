package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	N     *TreeNode
	level int
}

func main() {
	root := TreeNode{Val: 100}
	left := TreeNode{Val: 200}
	right := TreeNode{Val: 300}
	root.Left = &left
	root.Right = &right

	result := levelOrder(&root)
	fmt.Println(result)
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	// 任务队列
	queue := make([]Node, 0)
	// 初始化任务队列
	queue = append(queue, Node{root, 1})

	var top = 0

	for top < len(queue) {
		// 任务出列
		n := queue[top]
		top++
		// 将结点添加到 二维切片中,如果 当前层级大于二维切片的外层长度，则表示这个节点是该层级的第一个节点，需要初始化切片，否则直接添加
		if n.level > len(result) {
			array := make([]int, 0)
			array = append(array, n.N.Val)
			result = append(result, array)
		} else {
			result[n.level-1] = append(result[n.level-1], n.N.Val)
		}

		// 如果左子树不为空，并且将左子树添加到任务队列
		if n.N.Left != nil {
			queue = append(queue, Node{N: n.N.Left, level: n.level + 1})
		}

		// 如果右子树不为空，并且将右子树添加到任务队列
		if n.N.Right != nil {
			queue = append(queue, Node{N: n.N.Right, level: n.level + 1})
		}
	}

	return result
}
