package array_stack

import (
	"fmt"
	"testing"
)

func TestStack_Init(t *testing.T) {
	arrayStack := newArrayStack(10)
	arrayStack.Push(1)
	arrayStack.Push(2)
	arrayStack.Push(3)
	fmt.Print("---------------- \n")
	arrayStack.Show()
	arrayStack.Pop()
	arrayStack.Push(4)
	arrayStack.Push(5)
	arrayStack.Show()
	arrayStack.IsEmpty()
}
