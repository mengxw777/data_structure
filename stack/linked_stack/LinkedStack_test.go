package linked_stack

import (
	"fmt"
	"testing"
)

func TestNewLinkedStack(t *testing.T) {
	linkedStack := NewLinkedStack(10)
	linkedStack.Push(1)
	linkedStack.Push(2)
	linkedStack.Push(3)
	linkedStack.Show()
	fmt.Print("---------------- \n")
	linkedStack.Pop()
	linkedStack.Push(4)
	linkedStack.Push(5)
	linkedStack.Show()
}
