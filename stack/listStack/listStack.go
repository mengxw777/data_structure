package listStack

import (
	singleList "data_structure/linked_list/single_list"
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
	limit uint
	stack singleList.List
}

func (stack *Stack) Init(limit uint) {
	list := singleList.List{}
	list.Init()
	stack.limit = limit
	stack.mutex = new(sync.RWMutex)
	stack.stack = list
}

//	入栈
func (stack *Stack) Push(data interface{}) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if stack.stack.Size >= stack.limit {
		stack.stack.Delete(0)
	}

	stack.stack.Insert(0, &singleList.Node{Data: data})
}

//	出栈
func (stack *Stack) Pop() *singleList.Node {
	if stack.IsEmpty() {
		return nil
	}

	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	data := stack.stack.Get(0)
	stack.stack.Delete(0)
	return data
}

func (stack *Stack) IsEmpty() bool {
	stack.mutex.RLock()
	defer stack.mutex.RUnlock()

	return stack.stack.Size == 0
}

func (stack *Stack) Show() {
	stack.mutex.RLock()
	defer stack.mutex.RUnlock()

	head := stack.stack.Head
	var i uint
	for i = 0; i < stack.stack.Size; i++ {
		fmt.Printf("栈内数据: %v \n", head.Data)
		head = head.Next
	}

}
