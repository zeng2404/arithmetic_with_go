package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//firstList := ListNode{Val: 1}
	//firstList.Next = &ListNode{Val: 4, Next: &ListNode{Val: 3}}
	//secondList := ListNode{Val: 5}
	//secondList.Next = &ListNode{Val: 6, Next: &ListNode{Val: 4}}
	firstList := ListNode{Val: 1}
	secondList := ListNode{Val: 9}
	secondList.Next = &ListNode{Val: 9}
	resultList := addTwoNumbers(&firstList, &secondList)

	for resultList != nil {
		fmt.Println(resultList.Val)
		resultList = resultList.Next
	}
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	// 用于判断上一位是否需要进一位
//	lastVal := 0
//
//	// 返回的链表
//	var root *ListNode
//
//	var l3 *ListNode
//
//	// 当比较短的链表迭代结束以后,则退出循环
//	for l1 != nil && l2 != nil {
//		// 	如果等于 nil 则表示当前为第一个节点
//		if l3 == nil {
//			l3 = &ListNode{}
//			root = l3
//		} else {
//			// 如果不是则要新建下一个节点，并且完成拼接以及迭代到下一个节点
//			next := &ListNode{}
//			l3.Next = next
//			l3 = l3.Next
//		}
//		l3.Val = (l1.Val + l2.Val + lastVal) % 10
//		if l1.Val+l2.Val+lastVal >= 10 {
//			lastVal = 1
//		} else {
//			lastVal = 0
//		}
//		l1 = l1.Next
//		l2 = l2.Next
//	}
//
//	// 拼接剩下的部分
//	if l1 != nil {
//		for l1 != nil {
//			// 由于两个链表都是非空链表，直接新建一个节点，拼接到 l3
//			next := &ListNode{}
//			next.Val = (l1.Val + lastVal) % 10
//			if l1.Val+lastVal >= 10 {
//				lastVal = 1
//			} else {
//				lastVal = 0
//			}
//			l3.Next = next
//			l3 = l3.Next
//			l1 = l1.Next
//		}
//	}
//
//	if l2 != nil {
//		for l2 != nil {
//			next := &ListNode{}
//			next.Val = (l2.Val + lastVal) % 10
//			if l2.Val+lastVal >= 10 {
//				lastVal = 1
//			} else {
//				lastVal = 0
//			}
//			l3.Next = next
//			l3 = l3.Next
//			l2 = l2.Next
//		}
//	}
//
//	if l1 == nil && l2 == nil && lastVal > 0 {
//		// 两个链表同时走完，但是相加进了一位，则要创建一个新节点去存储
//		next := &ListNode{}
//		next.Val = lastVal
//		l3.Next = next
//		//lastVal = 0
//	}
//	return root
//}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	// 用于判断上一位是否需要进一位
//	lastVal := 0
//
//	// 返回的链表
//	var root *ListNode
//
//	var l3 *ListNode
//
//	// 当比较短的链表迭代结束以后,则退出循环
//	for l1 != nil || l2 != nil {
//		// 	如果等于 nil 则表示当前为第一个节点
//		if l3 == nil {
//			l3 = &ListNode{}
//			root = l3
//		} else {
//			// 如果不是则要新建下一个节点，并且完成拼接以及迭代到下一个节点
//			next := &ListNode{}
//			l3.Next = next
//			l3 = l3.Next
//		}
//
//		if l1 != nil && l2 != nil {
//			l3.Val = (l1.Val + l2.Val + lastVal) % 10
//			if l1.Val+l2.Val+lastVal >= 10 {
//				lastVal = 1
//			} else {
//				lastVal = 0
//			}
//			l1 = l1.Next
//			l2 = l2.Next
//		} else if l1 != nil {
//			l3.Val = (l1.Val + lastVal) % 10
//			if l1.Val + lastVal >= 10 {
//				lastVal = 1
//			} else {
//				lastVal = 0
//			}
//			l1 = l1.Next
//		} else if l2 != nil {
//			l3.Val = (l2.Val + lastVal) % 10
//			if l2.Val + lastVal >= 10 {
//				lastVal = 1
//			} else {
//				lastVal = 0
//			}
//			l2 = l2.Next
//		}
//	}
//
//	if l1 == nil && l2 == nil && lastVal > 0 {
//		// 两个链表同时走完，但是相加进了一位，则要创建一个新节点去存储
//		next := &ListNode{}
//		next.Val = lastVal
//		l3.Next = next
//		//lastVal = 0
//	}
//	return root
//}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	// 用于判断上一位是否需要进一位
//	lastVal := 0
//
//	// 返回的链表
//	var root *ListNode
//
//	var l3 *ListNode
//
//	// 当比较短的链表迭代结束以后,则退出循环
//	for l1 != nil || l2 != nil {
//		// 	如果等于 nil 则表示当前为第一个节点
//		if l3 == nil {
//			l3 = &ListNode{}
//			root = l3
//		} else {
//			// 如果不是则要新建下一个节点，并且完成拼接以及迭代到下一个节点
//			next := &ListNode{}
//			l3.Next = next
//			l3 = l3.Next
//		}
//
//		if l1 != nil && l2 != nil {
//			l3.Val = (l1.Val + l2.Val + lastVal) % 10
//			lastVal = (l1.Val + l2.Val + lastVal) / 10
//			l1 = l1.Next
//			l2 = l2.Next
//		} else if l1 != nil {
//			l3.Val = (l1.Val + lastVal) % 10
//			lastVal = (l1.Val + lastVal) / 10
//			l1 = l1.Next
//		} else if l2 != nil {
//			l3.Val = (l2.Val + lastVal) % 10
//			lastVal = (l2.Val + lastVal) / 10
//			l2 = l2.Next
//		}
//	}
//
//	if l1 == nil && l2 == nil && lastVal > 0 {
//		// 两个链表同时走完，但是相加进了一位，则要创建一个新节点去存储
//		next := &ListNode{}
//		next.Val = lastVal
//		l3.Next = next
//	}
//	return root
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 用于判断上一位是否需要进一位
	lastVal := 0

	// 返回的链表
	var root *ListNode

	var l3 *ListNode

	// 当比较短的链表迭代结束以后,则退出循环
	for l1 != nil && l2 != nil {
		// 	如果等于 nil 则表示当前为第一个节点
		if l3 == nil {
			l3 = &ListNode{}
			root = l3
		} else {
			// 如果不是则要新建下一个节点，并且完成拼接以及迭代到下一个节点
			next := &ListNode{}
			l3.Next = next
			l3 = l3.Next
		}
		l3.Val = (l1.Val + l2.Val + lastVal) % 10
		lastVal = (l1.Val + l2.Val + lastVal) / 10
		l1 = l1.Next
		l2 = l2.Next
	}

	// 拼接剩下的部分
	if l1 != nil {
		for l1 != nil {
			// 由于两个链表都是非空链表，直接新建一个节点，拼接到 l3
			next := &ListNode{}
			next.Val = (l1.Val + lastVal) % 10
			lastVal = (l1.Val + lastVal) / 10
			l3.Next = next
			l3 = l3.Next
			l1 = l1.Next
		}
	}

	if l2 != nil {
		for l2 != nil {
			next := &ListNode{}
			next.Val = (l2.Val + lastVal) % 10
			lastVal = (l2.Val + lastVal) / 10
			l3.Next = next
			l3 = l3.Next
			l2 = l2.Next
		}
	}

	if l1 == nil && l2 == nil && lastVal > 0 {
		// 两个链表同时走完，但是相加进了一位，则要创建一个新节点去存储
		next := &ListNode{}
		next.Val = lastVal
		l3.Next = next
	}
	return root
}
