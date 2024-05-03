package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			head.Next = list2
			break
		} else if list2 == nil {
			head.Next = list1
			break
		} else if list2.Val >= list1.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	return res.Next
}
func PrintListNode(l1 *ListNode) {
	next := l1
	ar := make([]int, 0)
	for next != nil {
		ar = append(ar, next.Val)

		next = next.Next
	}
	fmt.Println(ar)
}
