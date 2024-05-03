package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	res := isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
	)
	fmt.Println(res)
	res = isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil},
		&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}},
	)
	fmt.Println(res)
	res = isSameTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 1, Left: nil, Right: nil}},
		&TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 2, Left: nil, Right: nil}},
	)
	fmt.Println(res)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
