// Реализовать пересечение двух неупорядоченных множеств
package main

import "fmt"

// func intersection(setA, setB []int) []int {
// 	intersection := make([]int, 0)

// 	setAMap := make(map[int]bool)
// 	for _, elem := range setA {
// 		setAMap[elem] = true
// 	}

// 	for _, elem := range setB {
// 		if setAMap[elem] {
// 			intersection = append(intersection, elem)
// 		}
// 	}
// 	return intersection
// }

// func main() {
// 	setA := []int{1, 2, 3, 4, 5}
// 	setB := []int{4, 5, 6, 7, 8}

// 	intersection := intersection(setA, setB)

// 	fmt.Println(intersection)

// }
func main() {
	setA := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
	}
	setB := map[int]struct{}{
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
	}

	intersection := make(map[int]struct{})

	for element := range setA {
		if _, ok := setB[element]; ok {
			intersection[element] = struct{}{}
		}
	}

	fmt.Println(intersection)
}
