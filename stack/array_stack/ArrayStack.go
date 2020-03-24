package array_stack

import (
	"fmt"
	"sync"
)

type StackInterface interface {
	Push()
	Pop()
	IsEmpty()
}

type Stack struct {
	mutex *sync.RWMutex
	limit int
	stack []interface{}
}

//	初始化
func newArrayStack(limit int) *Stack {
	return &Stack{
		mutex: new(sync.RWMutex),
		limit: limit,
		stack: nil,
	}
}

//	入栈
func (stack *Stack) Push(data interface{}) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.stack) >= stack.limit {
		stack.Pop()
	}

	tmp := append(stack.stack, 0)
	copy(tmp[0+1:], tmp[0:])
	tmp[0] = data
	stack.stack = tmp
}

//	出栈
func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}

	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	item := stack.stack[0]
	stack.stack = stack.stack[1:]
	return item
}

//	判断是否为空
func (stack *Stack) IsEmpty() bool {
	stack.mutex.RLock()
	defer stack.mutex.RUnlock()

	return len(stack.stack) == 0
}

func (stack *Stack) Show() {
	stack.mutex.RLock()
	defer stack.mutex.RUnlock()

	fmt.Printf("栈内数据: %v \n", stack.stack)
}
