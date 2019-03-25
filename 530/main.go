package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	var root *TreeNode = new(TreeNode)
	root.Val = 1

	var right *TreeNode = new(TreeNode)
	right.Val = 5
	root.Right = right

	var left2 *TreeNode = new(TreeNode)
	left2.Val = 3
	root.Right.Left = left2

	a := getMinimumDifference(root)
	fmt.Println(a)
}

func getMinimumDifference(root *TreeNode) int {
	r := trans(root)
	return r
}

var result int = int(^uint(0) >> 1)
var preNode *TreeNode

//遍历二叉树
func trans(root *TreeNode) int {

	if root == nil {
		return 0
	}
	trans(root.Left)

	if preNode != nil {
		if (root.Val - preNode.Val) < result {
			result = root.Val - preNode.Val
		}
	}
	preNode = root

	trans(root.Right)
	//fmt.Println(root) //后序遍历
	return result
}
