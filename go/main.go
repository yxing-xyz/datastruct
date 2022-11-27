package main

func main() {
	node1 := NewBinaryTreeNode(1)
	node2 := NewBinaryTreeNode(2)
	node3 := NewBinaryTreeNode(3)
	node4 := NewBinaryTreeNode(4)
	node5 := NewBinaryTreeNode(5)
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node4.Left = node5

	node1.MidOrder()
}

//1,2,4,5,3
// 54213

// 54231
