package bst

import (
	"fmt"
)

type Tree struct {
	root *TreeNode
}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *Tree) insert(val int) {
	if t.root == nil {
		t.root = &TreeNode{}
		t.root.val = val
	} else {
		current := t.root
		for {
			if val > current.val {
				if current.right == nil {
					current.right = &TreeNode{}
					current.right.val = val
					break
				}
				current = current.right
			} else {
				if current.left == nil {
					current.left = &TreeNode{}
					current.left.val = val
					break
				}
				current = current.left
			}
		}
	}
}

func (t *Tree) search(val int) bool {
	return search(val, t.root)
}

func (t *Tree) printInorder() {
	fmt.Printf("Inorder is as follows :")
	printInorder(t.root)
	fmt.Println()
}

func (t *Tree) printPreOrder() {
	fmt.Printf("Preorder is as follows :")
	printPreOrder(t.root)
	fmt.Println()
}

func (t *Tree) printPostOrder() {
	fmt.Printf("PostOrder is as follows :")
	printPostOrder(t.root)
	fmt.Println()
}

func (t *Tree) getHeight() int {
	return getHeight(t.root)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.left)
	rightHeight := getHeight(root.right)
	maxHeight := leftHeight
	if rightHeight > leftHeight {
		maxHeight = rightHeight
	}
	return maxHeight + 1
}

func printPostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printNodeValue(root)
	printPostOrder(root.left)
	printPostOrder(root.right)
}

func printPreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printNodeValue(root)
	printPreOrder(root.left)
	printPreOrder(root.right)
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.left)
	printNodeValue(root)
	printInorder(root.right)
}

func search(val int, root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.val == val {
		return true
	} else if val > root.val {
		return search(val, root.right)
	} else {
		return search(val, root.left)
	}
}

func printNodeValue(root *TreeNode) {
	fmt.Printf("%d ", root.val)
}

func TestBST() {
	tree := Tree{}
	tree.insert(2)
	tree.insert(1)
	tree.insert(-1)
	tree.insert(3)
	tree.insert(4)

	tree.printInorder()
	tree.printPreOrder()
	tree.printPostOrder()

	fmt.Println(tree.search(5))
	fmt.Println(tree.search(4))
	fmt.Println(tree.search(-1))
	fmt.Println(tree.search(-2))

	fmt.Printf("Height of tree is %d", tree.getHeight())
}
