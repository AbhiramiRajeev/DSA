package main

import "fmt"

/*
Input: arr = [1, 2, 3, 4, 6], target = 6
Output: [1, 3]  // 2 + 4 = 6
Approach (Two Pointers):

Have left = 0, right = len(arr)-1

Calculate current_sum = arr[left] + arr[right]

If it equals the target → return indices

If it’s smaller → move left++ (need a bigger sum)

If it’s larger → move right-- (need a smaller sum)
*/

func main() {
	fmt.Println(pairWithTargetSum([]int{1, 2, 3, 4, 6}, 6))
}

func pairWithTargetSum(arr []int, target int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		if arr[left]+arr[right] == target {
			return []int{left, right}
		} else if arr[left]+arr[right] < target {
			left++
		} else {
			right--
		}

	}
	return []int{}
}
