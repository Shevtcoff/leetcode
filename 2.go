package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/* func main() {

	res := addTwoNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
	)
	printListNode(res)

	res = addTwoNumbers(
		&ListNode{Val: 0, Next: nil},
		&ListNode{Val: 0, Next: nil},
	)
	printListNode(res)

	res = addTwoNumbers(
		&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 3, Next: &ListNode{Val: 9, Next: nil}}}},
		&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 8, Next: nil}}},
	)
	printListNode(res)
} */

/*
Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

*/

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	shift := 0

	for l1 != nil || l2 != nil {

		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := shift + val1 + val2
		shift = sum / 10

		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
	}

	if shift > 0 {
		cur.Next = &ListNode{1, nil}
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
