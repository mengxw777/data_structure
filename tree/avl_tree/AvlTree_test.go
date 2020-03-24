package avl_tree

import (
	"fmt"
	"testing"
)

func addData() *AVLTreeNode {
	data := []int{80, 85, 60, 82, 90, 87, 100}
	root := NewAVLTree(70)
	for _, value := range data {
		root = root.Insert(value)
	}
	return root
}

func TestAVLTreeNode_Search(t *testing.T) {
	tree := addData()
	fmt.Printf("寻找：90 结果：%v \n", tree.Search(90))
}

func TestAVLTreeNode_Max(t *testing.T) {
	tree := addData()
	fmt.Printf("最大 %v \n", tree.Max())
}

func TestAVLTreeNode_Min(t *testing.T) {
	tree := addData()
	fmt.Printf("最小 %v \n", tree.Min())
}

func TestAVLTreeNode_Delete(t *testing.T) {
	tree := addData()
	tree.DisplayNodesInOrder()
	fmt.Print("\n")
	tree.Delete(90)
	tree.DisplayNodesInOrder()
}
