package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//创建二叉树
	var root *TreeNode = new(TreeNode)
	root.Val = 1
	var left *TreeNode = new(TreeNode)
	left.Val = 2
	root.Left = left
	var right *TreeNode = new(TreeNode)
	right.Val = 4
	root.Right = right
	var left2 *TreeNode = new(TreeNode)
	left2.Val = 3
	root.Left.Left = left2
	//trans(root)
	a := isSameTree(root, root)
	fmt.Println(a)
}

//遍历二叉树
func trans(root *TreeNode) {
	if root == nil {
		return
	}
	//fmt.Println(root) //前序遍历
	trans(root.Left)
	fmt.Println(root.Val) //中序遍历
	trans(root.Right)
	//fmt.Println(root) //后序遍历
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	Ip := iterator(p)
	Iq := iterator(q)
	for vp := range Ip {
		vq := <-Iq
		if vp == nil && vq != nil {
			return false
		}

		if vp != nil && vq == nil {
			return false
		}

		if vp != nil && vq != nil {
			if vp != vq {
				return false
			}
		}
	}

	return true
}

// iterator（迭代器）, iterable（可迭代的数据）, generator（生成器）
// generator 是一种特殊的 iterable
// iterator 是用来 iterate 一个 iterable 的工具

func iterator(n *TreeNode) chan *TreeNode {
	c := make(chan *TreeNode)
	go func() {
		c <- n
		if n != nil {
			for n := range iterator(n.Left) {
				c <- n
			}
			for n := range iterator(n.Right) {
				c <- n
			}
		}
		close(c)
	}()
	return c
}
