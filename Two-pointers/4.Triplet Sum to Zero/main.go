package main

import (
	"fmt"
	"sort"
)

/*
Given an array of integers, find all unique triplets (three numbers) whose sum is zero.
Input:  [-3, 0, 1, 2, -1, 1, -2]
Output: [[-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]]

💡 Idea — How two pointers come into play

Sort the array first.
Sorting allows us to use the two-pointer technique efficiently.

Fix one number (say, at index i) → we now need a pair of numbers that sum to -arr[i].

Use two pointers (left, right) to find that pair:

If the sum of those two numbers is less than target → move left++

If greater → move right--

If equal → store the triplet, then move both pointers skipping duplicates
*/

func main() {
	fmt.Println(TripletSum([]int{-3, 0, 1, 2, -1, 1, -2}))
}

func TripletSum(arr []int) [][]int {
	result := [][]int{}
	sort.Ints(arr)

	for i := 0; i < len(arr)-2; i++ {
		//skip duplicates
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left, right := i+1, len(arr)-1

		for left < right {
			if arr[i]+arr[left]+arr[right] == 0 {
				result = append(result, []int{arr[i], arr[left], arr[right]})

				left++
				right--
				for left < right && arr[left] == arr[left-1] {
					left++
				}
				for left < right && arr[right] == arr[right+1] {
					right--
				}

			} else if arr[i]+arr[left]+arr[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result

}
