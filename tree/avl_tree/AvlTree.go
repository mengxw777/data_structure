package avl_tree

import (
	"fmt"
	"math"
)

type AVLTreeNode struct {
	key         int
	height      int
	left, right *AVLTreeNode
}

func NewAVLTree(key int) *AVLTreeNode {
	return &AVLTreeNode{key: key}
}

//	插入
func (node *AVLTreeNode) Insert(key int) *AVLTreeNode {
	if node == nil {
		return NewAVLTree(key)
	}

	if key < node.key {
		node.left = node.left.Insert(key)
	}

	if key > node.key {
		node.right = node.right.Insert(key)

	}

	node = node.rebalanced()
	node.recalculateHeight()
	return node
}

//	左旋
func (node *AVLTreeNode) leftRotate() *AVLTreeNode {
	newRoot := node.right
	node.right = newRoot.left
	newRoot.left = node

	node.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

//	右旋
func (node *AVLTreeNode) rightRotate() *AVLTreeNode {
	newRoot := node.left
	node.left = newRoot.right
	newRoot.right = node

	node.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

//	计算高度
func (node *AVLTreeNode) recalculateHeight() {
	node.height = int(math.Max(float64(node.left.getHeight()), float64(node.right.getHeight()))) + 1
}

//	平衡节点
func (node *AVLTreeNode) rebalanced() *AVLTreeNode {
	balanceFactor := node.left.getBalancedFactor() - node.right.getBalancedFactor()

	if balanceFactor == -2 {
		if node.right.left.getHeight() > node.right.right.getHeight() {
			node.right = node.right.rightRotate()
		}
		return node.leftRotate()
	} else if balanceFactor == 2 {
		if node.left.right.getHeight() > node.left.left.getHeight() {
			node.left = node.left.leftRotate()
		}
		return node.rightRotate()
	}

	return node
}

//	获取节点高度
func (node *AVLTreeNode) getHeight() int {
	if node != nil {
		return node.height
	}

	return -1
}

//	获取平衡因子
func (node *AVLTreeNode) getBalancedFactor() int {
	if node != nil {
		return node.left.getHeight() - node.right.getHeight()
	}

	return -1
}

//	最大节点
func (node *AVLTreeNode) Max() *AVLTreeNode {
	if node != nil && node.right != nil {
		return node.right.Max()
	}

	return node
}

//	最小节点
func (node *AVLTreeNode) Min() *AVLTreeNode {
	if node != nil && node.left != nil {
		return node.left.Min()
	}

	return node
}

//	寻找节点
func (node *AVLTreeNode) Search(key int) *AVLTreeNode {
	if node == nil {
		return nil
	}

	if key > node.key {
		return node.right.Search(key)
	}

	if key > node.key {
		return node.left.Search(key)
	}

	return node
}

//	删除节点
func (node *AVLTreeNode) Delete(key int) *AVLTreeNode {
	if node == nil {
		return node
	}

	if key < node.key {
		node.left = node.left.Delete(key)
	} else if key > node.key {
		node.right = node.right.Delete(key)
	} else {
		if node.left != nil && node.right != nil {
			rightMinNode := node.right.Min()
			node.key = rightMinNode.key
			node.right = node.right.Delete(rightMinNode.key)
		} else if node.left != nil {
			node = node.left
		} else if node.right != nil {
			node = node.right
		} else {
			node = nil
			return node
		}
	}

	return node.rebalanced()
}

//	中序打印
func (node *AVLTreeNode) DisplayNodesInOrder() {
	if node.left != nil {
		node.left.DisplayNodesInOrder()
	}
	fmt.Print(node.key, " ")

	if node.right != nil {
		node.right.DisplayNodesInOrder()
	}
}
