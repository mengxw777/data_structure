package main

import (
	set "data_structure/set/set"
	"log"
)

func main() {

	set1 := set.New()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2 := set.New()
	set2.Add(1)
	set2.Add(2)
	set2.Add(3)
	set2.Add(4)
	set2.Add(5)
	set2.Add(6)

	//log.Printf("测试包含 %v", set1.Contains(1))
	//log.Printf("测试移除 %v %v", set1.Remove(2), set1)
	//log.Printf("测试交集 %v", set1.Intersect(set2))
	//log.Printf("测试差集 %v", set1.Diff(set2))
	//log.Printf("测试并集 %v", set1.Union(set2))
	log.Printf("测试子集 %v", set1.IsSubsetOf(set2))

}
