package main

import (
	"data_structure/hashMap/hashMap"
	"log"
)

func main() {
	newHashMap := hashMap.New()
	newHashMap.Add("1", 1)
	newHashMap.Add("2", 2)
	newHashMap.Add("3", 3)
	newHashMap.Remove("1")

	log.Print(newHashMap.Get("1"))
	log.Print(newHashMap.Get("2"))
}
