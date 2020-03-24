package binary_search_tree

import (
	"fmt"
)

type treeInterface interface {
	Insert(value int)
	Remove()
	Minimum()
	Maximum()
	Successor()
	Search()
}

type tree struct {
	value               int
	parent, left, right *tree
}

func New() *tree {
	return &tree{}
}

//	插入
func Insert(tree *tree, value int) {
	//	根节点
	if tree.value == 0 {
		tree.value = value
		return
	}

	//	插入之小于等于当前根节点值，向左插入
	if value <= tree.value {
		if tree.left == nil {
			tree.left = New()
			tree.left.parent = tree
			tree.left.value = value
			return
		} else {
			Insert(tree.left, value)
		}
	}

	//	插入之大于等于当前根节点值，向右插入
	if value >= tree.value {
		if tree.right == nil {
			tree.right = New()
			tree.right.parent = tree
			tree.right.value = value
			return
		} else {
			Insert(tree.right, value)
		}
	}
}

func Search(tree *tree, value int) *tree {
	if tree == nil || tree.value == value {
		return tree
	}

	if value < tree.value {
		return Search(tree.left, value)
	} else {
		return Search(tree.right, value)
	}
}

//	最小值
func Minimum(tree *tree) *tree {
	if tree.value == 0 {
		return nil
	}

	for tree.left != nil {
		tree = tree.left
	}

	return tree
}

//	最大值
func Maximum(tree *tree) *tree {
	if tree.value == 0 {
		return nil
	}

	for tree.right != nil {
		tree = tree.right
	}

	return tree
}

//	前驱节点：对一颗二叉树进行中序遍历，按照遍历后的顺序，当前节点的前一个节点为该节点的前驱节点
//	后继节点：对一颗二叉树进行中序遍历，按照遍历后的顺序，当前节点的后一个节点为该节点的后继节点

func Precursor(tree *tree) *tree {
	if tree.left != nil {
		return Maximum(tree.left)
	}

	parent := tree.parent
	for parent != nil && tree == parent.left {
		tree = parent
		parent = tree.parent
	}

	return parent
}

//	后继
func Successor(x *tree) *tree {
	if x.right != nil {
		//	右子树不为空，它的后继为右子树的最小值
		return Minimum(x.right)
	}

	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}

	return y
}

func Show(tree *tree) {
	if tree.value == 0 {
		return
	}

	var tmp int
	if tree.parent == nil {
		tmp = 0
	} else {
		tmp = tree.parent.value
	}
	fmt.Printf("当前节点是: %v 父节点是: %v \n", tree.value, tmp)

	if tree.left != nil {
		Show(tree.left)
	}

	if tree.right != nil {
		Show(tree.right)
	}
}

func InOrderTraversal(tree *tree) {
	if tree.value == 0 {
		return
	}

	if tree.left != nil {
		InOrderTraversal(tree.left)
	}

	fmt.Printf("当前节点是: %v \n", tree.value)

	if tree.right != nil {
		InOrderTraversal(tree.right)
	}
}
