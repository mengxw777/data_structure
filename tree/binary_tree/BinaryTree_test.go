package binary_tree

import (
	"fmt"
	"testing"
)

func addData() {
	//	第一层
	A := NewBinaryTree()
	A.SetValue("A")

	//	第二层
	B := NewBinaryTree()
	B.SetValue("B")
	A.SetLeft(B)

	C := NewBinaryTree()
	C.SetValue("C")
	A.SetRight(C)

	//	第三层子树
	D := NewBinaryTree()
	D.SetValue("D")
	B.SetLeft(D)

	F := NewBinaryTree()
	F.SetValue("F")
	B.SetRight(F)

	//	第四层
	E := NewBinaryTree()
	E.SetValue("E")
	D.SetLeft(E)

	G := NewBinaryTree()
	G.SetValue("G")
	F.SetLeft(G)

	fmt.Print("前序 : \n")
	PreloaderTraversal(A)
	fmt.Print("中序 : \n")
	InOrderTraversal(A)
	fmt.Print("后序 : \n")
	PostorderTraversal(A)
}

func TestAdd(t *testing.T) {
	addData()
}
