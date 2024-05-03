package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	res := deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}})
	PrintListNode(res)
	res = deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}})
	PrintListNode(res)
	res = deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}})
	PrintListNode(res)
}

/*
Input: head = [1,1,2]
Output: [1,2]

Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/
func deleteDuplicates(head *ListNode) *ListNode {
	res := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}

	}
	return res
}

func deleteDuplicates2(head *ListNode) *ListNode {
	res := head
	if head == nil || head.Next == nil {
		return res
	}
	prev := head
	next := head.Next
	for head != nil && head.Next != nil {
		if prev.Val == next.Val {
			if next.Next == nil {
				prev.Next = nil
			} else {
				prev.Next = next.Next
			}
		} else {
			prev = next
		}
		next = next.Next
	}
	return res
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
