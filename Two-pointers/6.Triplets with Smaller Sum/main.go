package main

import (
	"fmt"
	"sort"
)

/*
Given an array of numbers arr and a target value target,
find the count of all triplets (i, j, k) such that:
arr[i] + arr[j] + arr[k] < target
Example

Input:
arr = [-1, 0, 2, 3], target = 3

Output:
2

Explanation:
The triplets are:

(-1, 0, 2) → sum = 1 < 3

(-1, 0, 3) → sum = 2 < 3
💡 Logic (Two Pointer Pattern)

Sort the array.

Fix one number arr[i].

Use two pointers left = i + 1, right = len(arr) - 1.

Compute sum = arr[i] + arr[left] + arr[right].

If sum < target, then all triplets between left and right are valid,
because the array is sorted → every element between left and right will make the sum smaller.
So add (right - left) to count, and move left++.

Else, move right-- to reduce sum.

Continue until left >= right.
*/

func main() {
	fmt.Println(TripletSmallerSum([]int{-1, 0, 2, 3}, 3))
}

func TripletSmallerSum(arr []int, target int) int {
	sort.Ints(arr)
	count := 0
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum < target {
				count += right - left
				left++
			} else {
				right--
			}
		}

	}
	return count
}
