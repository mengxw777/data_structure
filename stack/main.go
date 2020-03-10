package main

import (
	"data_structure/stack/arrayStack"
	"data_structure/stack/listStack"
	"fmt"
)

//	栈 	->	先进后出
//	队列	->	先进先出

func main() {
	//	数组实现栈
	stack := arrayStack.Stack{}
	stack.Init(10)

	fmt.Print("入栈 \n")
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(5)
	stack.Push(6)

	stack.Show()

	fmt.Print("出栈 \n")

	fmt.Print("出栈第一个 \n")
	stack.Pop()
	stack.Show()

	fmt.Print("出栈第二个 \n")
	stack.Pop()
	stack.Show()

	fmt.Print("出栈第三个 \n")
	stack.Pop()
	stack.Show()

	fmt.Print("入栈一个 \n")
	stack.Push(100)

	stack.Show()

	//	链栈
	stack2 := listStack.Stack{}
	stack2.Init(10)

	stack2.Push(1)
	stack2.Push(2)
	stack2.Push(3)

	stack2.Show()

	stack2.Pop()
	stack2.Pop()

	stack2.Show()
}
