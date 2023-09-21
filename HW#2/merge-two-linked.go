package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var startHead *ListNode = &ListNode{}
	var preHead *ListNode = startHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			startHead.Next = list1
			list1 = list1.Next
		} else {
			startHead.Next = list2
			list2 = list2.Next
		}
		startHead = startHead.Next
	}

	if list1 == nil {
		startHead.Next = list2
	} else if list2 == nil {
		startHead.Next = list1
	}

	return preHead.Next
}

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}

	list2 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	merged := mergeTwoLists(list1, list2)

	for node := merged; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
