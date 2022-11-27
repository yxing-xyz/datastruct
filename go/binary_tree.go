package main

import "fmt"

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewBinaryTreeNode(data int) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data: data,
	}
}

func (root *BinaryTreeNode) GetTreeNodeNum() int {
	if root == nil {
		return 0
	}
	// 左字树计算 + 右子数计算 + 1(本身这个节点)
	return root.Left.GetTreeNodeNum() + root.Right.GetTreeNodeNum() + 1
}

func (root *BinaryTreeNode) GetTreeDegree() int {
	maxDeggree := 0
	if root == nil {
		return maxDeggree
	}

	leftDegree := root.Left.GetTreeDegree()
	rightDegree := root.Right.GetTreeDegree()
	if leftDegree > rightDegree {
		maxDeggree = leftDegree
	} else {
		maxDeggree = rightDegree
	}
	return leftDegree + 1
}

func (root *BinaryTreeNode) PreOrder() {
	if root != nil {
		fmt.Println(root.Data, " ")
		root.Left.PreOrder()
		root.Right.PreOrder()
	}
}

func (root *BinaryTreeNode) MidOrder() {
	if root != nil {
		root.Left.MidOrder()
		fmt.Println(root.Data, " ")
		root.Right.MidOrder()
	}
}

func (root *BinaryTreeNode) PostOrder() {
	if root != nil {
		root.Left.PostOrder()
		root.Right.PostOrder()
		fmt.Println(root.Data, " ")
	}
}
