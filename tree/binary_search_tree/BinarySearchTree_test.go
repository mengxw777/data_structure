package binary_search_tree

import (
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := addData()

	Show(tree)
}

func TestMinimum(t *testing.T) {
	tree := addData()
	log.Printf("当前最小值为: %v", Minimum(tree))
}

func TestMaximum(t *testing.T) {
	tree := addData()
	log.Printf("当前最大值为: %v", Maximum(tree))
}

func TestSearch(t *testing.T) {
	tree := addData()
	log.Printf("查找指定值3: %v", Search(tree, 3))
}

func TestPrecursor(t *testing.T) {
	tree := addData()
	InOrderTraversal(tree)
	log.Printf("2前序为: %v", Precursor(Search(tree, 2)))
}

func TestSuccessor(t *testing.T) {
	tree := addData()
	InOrderTraversal(tree)
	log.Printf("11后继为: %v", Successor(Search(tree, 11)))
}

func addData() *tree {
	tree := New()
	Insert(tree, 10)
	Insert(tree, 1)
	Insert(tree, 2)
	Insert(tree, 3)
	Insert(tree, 11)
	Insert(tree, 12)
	return tree
}
