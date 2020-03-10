package hashMap

import (
	singleList "data_structure/linked_list/single_list"
	"sync"
)

const bucketCount = 16

type hashMapInterface interface {
	Add(key interface{}, value interface{})
	Get(key interface{})
	Remove(key interface{})
	getKey() string
}

type hashMap struct {
	mutex   *sync.RWMutex
	Element [bucketCount]*singleList.List
}

type hashData struct {
	key   string
	value interface{}
}

func New() *hashMap {
	hashMap := &hashMap{
		mutex: new(sync.RWMutex),
	}

	for i := 0; i < bucketCount; i++ {
		list := singleList.List{}
		list.Init()
		hashMap.Element[i] = &list
	}

	return hashMap
}

func (hashMap *hashMap) Add(key string, value interface{}) {
	hashMap.mutex.Lock()
	defer hashMap.mutex.Unlock()

	hashKey := hashMap.hashKey(key)
	hashMap.Element[hashKey].Append(
		&singleList.Node{
			Data: hashData{key: key, value: value},
		},
	)
}

func (hashMap *hashMap) Get(key string) interface{} {
	hashMap.mutex.Lock()
	defer hashMap.mutex.Unlock()

	hashKey := hashMap.hashKey(key)
	exist := hashMap.Element[hashKey]

	if exist.Size == 0 {
		return nil
	}

	var i uint
	current := exist.Head
	for i = 0; i < exist.Size; i++ {
		//	由interface转为struct
		hashData, _ := current.Data.(hashData)

		if hashData.key == key {
			return current.Data
		}

		current = current.Next
	}

	return nil
}

func (hashMap *hashMap) Remove(key string) bool {
	hashMap.mutex.Lock()
	defer hashMap.mutex.Unlock()

	hashKey := hashMap.hashKey(key)
	exist := hashMap.Element[hashKey]

	if exist.Size == 0 {
		return true
	}

	var i uint
	current := exist.Head

	for i = 0; i < exist.Size; i++ {
		hashData, _ := current.Data.(hashData)

		if hashData.key == key {
			exist.Delete(i)
			return true
		}

		current = current.Next
	}

	return true
}

func (hashMap *hashMap) hashKey(key string) int {
	sum := 0

	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}

	return sum % bucketCount
}
