// Реализовать бинарный поиск встроенными методами языка
package main

import (
	"fmt"
)

func BinarySearch(item int, arr []int) (int, bool) {
	l := 0
	r := len(arr) - 1

	for l <= r {
		// mid := l + (r-l)/2
		mid := (l + r) / 2

		if arr[mid] == item {
			return mid, true
		} else if arr[mid] < item {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 0, false
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	item := 13
	if res, ok := BinarySearch(item, arr); ok {
		fmt.Println(res)
	} else {
		fmt.Println("no item")
	}
	// fmt.Println(BinarySearch(item, arr))
}
