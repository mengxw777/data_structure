package singleList

import (
	"fmt"
	"sync"
)

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	mutex *sync.RWMutex
	Head  *Node
	Tail  *Node
	Size  uint
}

//	初始化
func (list *List) Init() {
	list.Size = 0
	list.Head = nil
	list.Tail = nil
	list.mutex = new(sync.RWMutex)
}

//	插入至尾部
func (list *List) Append(node *Node) bool {
	if node == nil {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size = 1
		return true
	}

	tail := list.Tail
	tail.Next = node
	list.Tail = node
	list.Size += 1

	return true
}

//	指定位置插入
func (list *List) Insert(index uint, node *Node) bool {
	if node == nil {
		return false
	}

	if index > list.Size {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Size += 1
		return true
	}

	ptr := list.Head
	var i uint
	for i = 1; i <= index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next
	ptr.Next = node
	node.Next = next
	list.Size += 1
	return true
}

//	获取指定位置元素
func (list *List) Get(index uint) *Node {
	if list == nil || list.Size == 0 || index > list.Size-1 {
		return nil
	}

	list.mutex.RLock()
	defer list.mutex.RUnlock()

	node := list.Head
	var i uint
	for i = 1; i <= index; i++ {
		node = node.Next
	}

	return node
}

//	通过data获取指定元素
func (list *List) Find(data interface{}) (uint, *Node) {
	if list == nil || list.Size == 0 {
		return 0, nil
	}

	node := list.Head
	var i uint
	for i = 1; i <= list.Size-1; i++ {
		node = node.Next
		if node.Data == data {

			return i, node
		}
	}

	return 0, nil
}

//	删除指定位置元素
func (list *List) Delete(index uint) bool {
	if list == nil || list.Size == 0 || index > list.Size-1 {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		head := list.Head.Next
		list.Head = head
		if list.Size == 1 {
			list.Tail = nil
		}
		list.Size -= 1
		return true
	}

	ptr := list.Head
	var i uint
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next

	ptr.Next = next.Next
	if index == list.Size-1 {
		list.Tail = ptr
	}
	list.Size -= 1

	return true
}

func (list *List) Show() {
	if list == nil {
		fmt.Print("空链")
		return
	}

	list.mutex.RLock()
	defer list.mutex.RUnlock()

	fmt.Printf("链长度: %d \n", list.Size)
	ptr := list.Head
	var i uint
	for i = 0; i < list.Size; i++ {
		fmt.Printf("No %d data is %v | next is %v \n", i, ptr.Data, ptr.Next)
		ptr = ptr.Next
	}
}
