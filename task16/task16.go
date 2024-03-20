// Реализовать быструю сортировку массива (quicksort) встроенными методами языка
package main

import "fmt"

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := (low - 1)

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return (i + 1)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	n := len(arr)
	quickSort(arr, 0, n-1)
	fmt.Println("Sorted array: ", arr)
}
