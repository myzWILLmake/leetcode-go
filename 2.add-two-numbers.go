package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode used for leetcode
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var currNode *ListNode
	var isCarry int
	var num1 int
	var num2 int
	for l1 != nil || l2 != nil {
		tmpNode := ListNode{}
		if currNode == nil {
			res = &tmpNode
			currNode = res
		} else {
			currNode.Next = &tmpNode
			currNode = &tmpNode
		}

		if l1 == nil {
			num1 = 0
		} else {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			num2 = 0
		} else {
			num2 = l2.Val
			l2 = l2.Next
		}

		tmpNode.Val = num1 + num2 + isCarry

		if tmpNode.Val >= 10 {
			tmpNode.Val -= 10
			isCarry = 1
		} else {
			isCarry = 0
		}
	}

	if isCarry == 1 {
		tmpNode := ListNode{}
		tmpNode.Val = 1
		currNode.Next = &tmpNode
	}
	return res
}

func printFunc(l *ListNode) {
	iter := l
	for iter != nil {
		fmt.Print(iter.Val)
		iter = iter.Next
	}
}

// func main() {
// 	n1_1 := ListNode{}
// 	n1_2 := ListNode{}
// 	n1_3 := ListNode{}

// 	n1_1.Val = 2
// 	n1_2.Val = 4
// 	n1_3.Val = 3

// 	n1_1.Next = &n1_2
// 	n1_2.Next = &n1_3

// 	n2_1 := ListNode{}
// 	n2_2 := ListNode{}
// 	n2_3 := ListNode{}

// 	n2_1.Val = 5
// 	n2_2.Val = 6
// 	n2_3.Val = 4

// 	n2_1.Next = &n2_2
// 	n2_2.Next = &n2_3

// 	res := addTwoNumbers(&n1_1, &n2_1)
// 	printFunc(res)
// }
