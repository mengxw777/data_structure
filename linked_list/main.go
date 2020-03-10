package main

import (
	singleList "data_structure/linked_list/single_list"
	"log"
)

//	缓存数量最大限制
const SIZE = 5

func main() {
	list := singleList.List{}
	list.Init()

	var i int
	for i = 0; i < 10; i++ {
		handler(i, &list)
	}

	for i = 5; i <= 10; i++ {
		handler(i, &list)
	}

	log.Print("打印现有缓存!")
	list.Show()

}

func handler(data interface{}, list *singleList.List) {
	index, node := list.Find(data)
	//	已缓存
	if node != nil {
		list.Insert(0, node)
		list.Delete(index)
		log.Printf("遇见已缓存数据，正在提权! %v", node.Data)

		return
	}

	//	未缓存
	if list.Size == SIZE {
		//	缓存溢出
		log.Printf("缓存溢出，删除末尾元素 %v，提权新元素! %v", list.Get(list.Size-1).Data, data)
		list.Delete(list.Size - 1)
		list.Insert(0, &singleList.Node{Data: data})
	} else {
		//	缓存未满
		list.Insert(0, &singleList.Node{Data: data})
		log.Printf("缓存未满，提权新元素! %v", data)
	}
}
