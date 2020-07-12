package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

func main() {
	three := TreeNode{Val: 3}
	one := TreeNode{Val: 1}
	two := TreeNode{Val: 2, Left: &one, Right: &three}
	seven := TreeNode{Val: 7}
	root := TreeNode{Val: 4, Left: &two, Right: &seven}
	a := insertIntoBST(&root, 5)
	Traverse(a)
}

func Traverse(n *TreeNode) {
	if n != nil {
		fmt.Println(n.Val)
		Traverse(n.Left)
		Traverse(n.Right)
	}
}
