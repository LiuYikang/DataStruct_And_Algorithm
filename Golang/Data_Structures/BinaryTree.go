package BinaryTree

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) Insert(val int) {
	if tree.root == nil {
		tree.root = &Node{val: val}
		return
	}

	currentNode := tree.root
	for {
		if val > currentNode.val {
			if currentNode.right == nil {
				currentNode.right = &Node{val: val}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &Node{val: val}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *BinaryTree) Search(val int) (*Node, bool) {
	if tree.root == nil {
		return nil, false
	}

	currentNode := tree.root
	for currentNode != nil {
		if val == currentNode.val {
			return currentNode, true
		} else if val < currentNode.val {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	return nil, false
}

func (tree *BinaryTree) InorderTraversal(subTree *Node) {
	if subTree.left != nil {
		tree.InorderTraversal(subTree.left)
	}
	fmt.Print(subTree.val)
	if subTree.right != nil {
		tree.InorderTraversal(subTree.right)
	}
}

func (tree *BinaryTree) PreorderTraversal(subTree *Node) {
	fmt.Print(subTree.val)
	if subTree.left != nil {
		tree.PreorderTraversal(subTree.left)
	}
	if subTree.right != nil {
		tree.PreorderTraversal(subTree.right)
	}
}

func (tree *BinaryTree) PostorderTraversal(subTree *Node) {
	if subTree.left != nil {
		tree.PostorderTraversal(subTree.left)
	}
	if subTree.right != nil {
		tree.PostorderTraversal(subTree.right)
	}
	fmt.Print(subTree.val)
}
