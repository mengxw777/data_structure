package set

import (
	"sync"
)

type setInterface interface {
	Add(element interface{}) bool
	Remove(element interface{}) bool
	Clear()
	Contains(element interface{}) bool
	Length() int
	Intersect(other *Set) *Set
	Diff(other *Set) *Set
	Union(other *Set) *Set
	IsSubsetOf(other *Set) bool
}

type Set struct {
	mutex    *sync.RWMutex
	Elements map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		mutex:    new(sync.RWMutex),
		Elements: make(map[interface{}]bool),
	}
}

func (set *Set) Add(element interface{}) bool {
	if set.Elements[element] == true {
		return false
	}

	set.mutex.Lock()
	defer set.mutex.Unlock()

	set.Elements[element] = true

	return true
}

func (set *Set) Remove(element interface{}) bool {
	if set.Elements[element] == false {
		return false
	}

	set.mutex.Lock()
	defer set.mutex.Unlock()

	delete(set.Elements, element)
	return true
}

func (set *Set) Clear() {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	set.Elements = map[interface{}]bool{}
}

func (set *Set) Contains(element interface{}) bool {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	return set.Elements[element]
}

func (set *Set) Length() int {
	set.mutex.RLock()
	defer set.mutex.RUnlock()

	return len(set.Elements)
}

//	交集	两个集合中相同的元素组合成的集合
func (set *Set) Intersect(other *Set) *Set {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	newSet := NewSet()

	for element := range set.Elements {
		if other.Contains(element) == true {
			newSet.Add(element)
		}
	}

	return newSet
}

//	差集	两个集合除相同元素外剩下元素的集合
func (set *Set) Diff(other *Set) *Set {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	newSet := NewSet()

	for setElement := range set.Elements {
		if other.Contains(setElement) == false {
			newSet.Add(setElement)
		}
	}

	return newSet
}

//	并集
func (set *Set) Union(other *Set) *Set {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	newSet := NewSet()

	for setElement := range set.Elements {
		newSet.Add(setElement)
	}

	for otherElement := range other.Elements {
		newSet.Add(otherElement)
	}

	return newSet
}

//	是否为子集
func (set *Set) IsSubsetOf(other *Set) bool {
	set.mutex.Lock()
	defer set.mutex.Unlock()

	for setElement := range set.Elements {
		if other.Contains(setElement) == false {
			return false
		}
	}

	return true
}
