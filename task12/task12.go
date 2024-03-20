// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество
package main

import "fmt"

func main() {
	line := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, l := range line {
		set[l] = struct{}{}
	}

	fmt.Println(set)
}
