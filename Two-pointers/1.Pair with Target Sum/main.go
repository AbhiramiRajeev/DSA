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
	fmt.Println(TargetSum([]int{2, 1, 3, 6, 4, 3, 5}, 6))
}

func pairWithTargetSumSorted(arr []int, target int) []int {
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

/*
Given a sorted array of integers arr and a target sum target, find all pairs of numbers whose sum equals the target.
*/

func TargetSum(arr []int, target int) [][]int {
	result := [][]int{}
	seen_map := make(map[int]bool)

	for _, val := range arr {
		num := target - val
		if seen_map[num] {
			result = append(result, []int{val, num})
		}
		seen_map[val] = true

	}
	return result
}
