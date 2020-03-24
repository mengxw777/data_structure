package singleList

import (
	"log"
	"testing"
)

func TestList_Append(t *testing.T) {
	list := NewSingleList()
	for i := 0; i < 10; i++ {
		handler(i, list)
	}
}

func handler(data interface{}, list *List) {
	index, node := list.Find(data)
	//	已缓存
	if node != nil {
		list.Insert(0, node)
		list.Delete(index)
		log.Printf("遇见已缓存数据，正在提权! %v", node.Data)

		return
	}

	//	未缓存
	if list.Size == 5 {
		//	缓存溢出
		log.Printf("缓存溢出，删除末尾元素 %v，提权新元素! %v", list.Get(list.Size-1).Data, data)
		list.Delete(list.Size - 1)
		list.Insert(0, &Node{Data: data})
	} else {
		//	缓存未满
		list.Insert(0, &Node{Data: data})
		log.Printf("缓存未满，提权新元素! %v", data)
	}
}
