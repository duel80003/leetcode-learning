package main

import (
	"fmt"
	"os"
)

// MyHashMap implement hash map
type MyHashMap struct {
	hashMap map[int]int
}

// Constructor intialize your data structure here.
func Constructor() MyHashMap {
	m := make(map[int]int)
	return MyHashMap{m}
}

// Put value will always be non-negative.
func (h *MyHashMap) Put(key int, value int) {
	h.hashMap[key] = value
}

// Get returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key
func (h *MyHashMap) Get(key int) int {
	v, ok := h.hashMap[key]
	if ok {
		return v
	}
	return -1
}

// Remove the mapping of the specified value key if this map contains a mapping for the key
func (h *MyHashMap) Remove(key int) {
	delete(h.hashMap, key)
}

func main() {
	t, _ := os.Getwd()
	fmt.Println(t)
}
