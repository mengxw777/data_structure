package set

import (
	"fmt"
	"testing"
)

func addData() *Set {
	newSet := NewSet()
	newSet.Add(1)
	newSet.Add(2)
	newSet.Add(3)
	newSet.Add(4)
	newSet.Add(5)
	return newSet
}

func addDat2() *Set {
	newSet := NewSet()
	newSet.Add(5)
	newSet.Add(11)
	newSet.Add(12)
	newSet.Add(13)
	newSet.Add(14)
	return newSet
}

func TestSet_Remove(t *testing.T) {
	newSet := addData()
	newSet.Remove(1)
	fmt.Printf("移除1, %v", newSet.Elements)
}

func TestSet_Contains(t *testing.T) {
	newSet := addData()
	fmt.Printf("是否包含5 %v \n", newSet.Contains(5))
	fmt.Printf("是否包含10 %v \n", newSet.Contains(5))
}

func TestSet_Length(t *testing.T) {
	newSet := addData()
	fmt.Printf("集合长度 %v \n", newSet.Length())
}

func TestSet_Intersect(t *testing.T) {
	newSet := addData()
	newSet2 := addDat2()
	fmt.Printf("集合交集 %v \n", newSet.Intersect(newSet2).Elements)
}

func TestSet_Diff(t *testing.T) {
	newSet := addData()
	newSet2 := addDat2()
	fmt.Printf("集合差集 %v \n", newSet.Diff(newSet2).Elements)
}

func TestSet_Union(t *testing.T) {
	newSet := addData()
	newSet2 := addDat2()
	fmt.Printf("集合并集 %v \n", newSet.Union(newSet2).Elements)
}

func TestSet_IsSubsetOf(t *testing.T) {
	newSet := addData()
	newSet2 := addDat2()
	fmt.Printf("是否为子集 %v \n", newSet.IsSubsetOf(newSet2))
}
