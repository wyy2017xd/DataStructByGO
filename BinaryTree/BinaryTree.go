package BinaryTree

import "fmt"

type BinaryNode struct {
	Data string
	LChild *BinaryNode
	RChild *BinaryNode
}

//先序(递归)
func (root *BinaryNode) PreOrder() {
	if (root != nil) {
		fmt.Println(root.Data)
		root.LChild.PreOrder()
		root.RChild.PreOrder()
	}
}
