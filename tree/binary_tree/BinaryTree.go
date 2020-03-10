package binary_tree

import "fmt"

type Tree struct {
	Value       interface{}
	left, right *Tree
}

func New() *Tree {
	return &Tree{}
}

func (tree *Tree) SetValue(value interface{}) {
	tree.Value = value
}

func (tree *Tree) SetLeft(node *Tree) {
	tree.left = node
}

func (tree *Tree) SetRight(node *Tree) {
	tree.right = node
}

//	前序遍历 ->	首先访问根，再先序遍历左（右）子树，最后先序遍历右（左）子树
func PreorderTraversal(tree *Tree) {
	if tree.Value == nil {
		return
	}

	fmt.Printf("当前节点是: %v \n", tree.Value)

	if tree.left != nil {
		PreorderTraversal(tree.left)
	}

	if tree.right != nil {
		PreorderTraversal(tree.right)
	}
}

//	中序遍历	->	首先中序遍历左（右）子树，再访问根，最后中序遍历右（左）子树
func InOrderTraversal(tree *Tree) {
	if tree.Value == nil {
		return
	}

	if tree.left != nil {
		InOrderTraversal(tree.left)
	}

	fmt.Printf("当前节点是: %v \n", tree.Value)

	if tree.right != nil {
		InOrderTraversal(tree.right)
	}
}

//	后序遍历	->	首先后序遍历左（右）子树，再后序遍历右（左）子树，最后访问根
func PostorderTraversal(tree *Tree) {
	if tree.Value == nil {
		fmt.Print("空树")
		return
	}

	if tree.left != nil {
		PostorderTraversal(tree.left)
	}

	if tree.right != nil {
		PostorderTraversal(tree.right)
	}

	fmt.Printf("当前节点是: %v \n", tree.Value)
}
