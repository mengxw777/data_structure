package hashMap

import (
	"fmt"
	"testing"
)

func addData() *hashMap {
	newMap := NewHashMap()
	newMap.Add("a", 1)
	newMap.Add("b", 2)
	newMap.Add("c", 3)
	return newMap
}

func TestHashMap_Get(t *testing.T) {
	newMap := NewHashMap()
	fmt.Printf("获取 2， %v", newMap)
}

func TestHashMap_Remove(t *testing.T) {
	newMap := NewHashMap()
	fmt.Printf("移除 2， %v", newMap.Remove("b"))
}
