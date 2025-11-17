package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"go", "python", "java", "c", "rust", "js"}
	fmt.Println(SortSring1(words))
}

//Sort strings by length (stable sort)

func SortSring1(arr []string) []string {
	sort.SliceStable(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})

	return arr
}

func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i; j < len(arr)-1; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}

		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr

}
