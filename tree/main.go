package main

import (
	"data_structure/tree/binary_tree"
	"fmt"
)

func main() {
	//	第一层
	A := binary_tree.New()
	A.SetValue("A")

	//	第二层
	B := binary_tree.New()
	B.SetValue("B")
	A.SetLeft(B)

	C := binary_tree.New()
	C.SetValue("C")
	A.SetRight(C)

	//	第三层子树
	D := binary_tree.New()
	D.SetValue("D")
	B.SetLeft(D)

	F := binary_tree.New()
	F.SetValue("F")
	B.SetRight(F)

	//	第四层
	E := binary_tree.New()
	E.SetValue("E")
	D.SetLeft(E)

	G := binary_tree.New()
	G.SetValue("G")
	F.SetLeft(G)

	fmt.Print("前序 : \n")
	binary_tree.PreorderTraversal(A)
	fmt.Print("中序 : \n")
	binary_tree.InOrderTraversal(A)
	fmt.Print("后序 : \n")
	binary_tree.PostorderTraversal(A)
}
