package main

import "fmt"

/*
Given a sorted array (which can have negative numbers), return a new array containing the squares of each number — also sorted in non-decreasing order.
Input:  [-3, -1, 0, 1, 2]
Output: [0, 1, 1, 4, 9]

If we square them directly:
[9, 1, 0, 1, 4] → not sorted anymore!
Because negative numbers (like -3) become large positive values after squaring.

Solution
Use two pointers:

left starts from beginning (most negative)

right starts from end (most positive)

compare abs(arr[left]) and abs(arr[right])

fill result array from the end (largest square first)

| left | right | arr[left] | arr[right] | bigger    | result      |
| ---- | ----- | --------- | ---------- | --------- | ----------- |
| 0    | 4     | -3        | 2          | 9 (left)  | [0,0,0,0,9] |
| 1    | 4     | -1        | 2          | 4 (right) | [0,0,0,4,9] |
| 1    | 3     | -1        | 1          | 1 (left)  | [0,0,1,4,9] |
| 2    | 3     | 0         | 1          | 1 (right) | [0,1,1,4,9] |


*/

func main() {
	fmt.Println(SquaringArray([]int{-3, -2, 1, 4, 5}))
}

func SquaringArray(arr []int) []int {
	left, right := 0, len(arr)-1
	lastIndex := len(arr) - 1
	result := make([]int, len(arr))

	for left <= right {
		leftSquare := arr[left] * arr[left]
		rightSquare := arr[right] * arr[right]

		if leftSquare > rightSquare {
			result[lastIndex] = leftSquare
			left++
		} else {
			result[lastIndex] = rightSquare
			right--
		}
		lastIndex--
	}
	return result
}
