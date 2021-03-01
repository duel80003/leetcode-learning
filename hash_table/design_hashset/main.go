package main

import "fmt"

// MyHashSet implement a hashset
type MyHashSet struct {
	hashSet map[int]int
}

// Constructor initialize your data structure here.
func Constructor() MyHashSet {
	var set MyHashSet
	set.hashSet = make(map[int]int)
	return set
}

// Add value to hashset
func (set *MyHashSet) Add(key int) {
	set.hashSet[key] = key
}

// Remove value from hashset
func (set *MyHashSet) Remove(key int) {
	delete(set.hashSet, key)

}

// Contains returns true if this set contains the specified element
func (set *MyHashSet) Contains(key int) bool {
	_, ok := set.hashSet[key]
	return ok
}

func main() {

	// set := Constructor()
	// set.Add(1)
	// set.Add(2)
	// fmt.Println(set.Contains(1))
	// fmt.Println(set.Contains(2))
	// set.Add(2)
	// fmt.Println(set.Contains(2))
	// set.Remove(2)
	// fmt.Println(set.Contains(2))

	// set2 := Constructor()
	// set2.Add(1)
	// set2.Remove(1)
	// set2.Remove(1)
	// fmt.Println(set2.Contains(1))

	set3 := Constructor()
	set3.Add(1)
	set3.Remove(2)
	// set3.Remove(1)
	fmt.Println(set3.Contains(1))
}
