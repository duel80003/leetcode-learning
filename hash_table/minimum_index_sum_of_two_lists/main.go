package main

import "fmt"

// func findRestaurant(list1 []string, list2 []string) []string {
// 	table := make(map[string]int)
// 	least := 2000
// 	for i := 0; i < len(list1); i++ {
// 		for j := 0; j < len(list2); j++ {
// 			if list1[i] == list2[j] {
// 				if least > (i + j) {
// 					least = i + j
// 					table[list1[i]] = i + j
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(least)
// 	r := []string{}
// 	for k, v := range table {
// 		if v <= least {
// 			r = append(r, k)
// 		}
// 	}
// 	return r
// }

func findRestaurant(list1 []string, list2 []string) []string {
	table := make(map[string]int)
	least := 2000
	r := []string{}
	for i, v := range list1 {
		table[v] = i
	}

	for i, v := range list2 {
		if val, ok := table[v]; ok {
			t := i + val
			if t < least {
				least = t
				r = []string{v}
			} else if t <= least {
				r = append(r, v)
			}
		}
	}
	return r
}

func main() {
	// [Shogun]
	a1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	a2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(a1, a2))

	// [Shogun]
	a3 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	a4 := []string{"KFC", "Shogun", "Burger King"}
	fmt.Println(findRestaurant(a3, a4))

	// ["KFC","Burger King","Tapioca Express","Shogun"]
	a5 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	a6 := []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	fmt.Println(findRestaurant(a5, a6))
}
