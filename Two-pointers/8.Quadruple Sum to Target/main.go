package main

import (
	"fmt"
	"sort"
)

/*
Given an array arr and a target number target, find all unique quadruplets [arr[a], arr[b], arr[c], arr[d]] such that:
Input: arr = [1, 0, -1, 0, -2, 2], target = 0
Output: [
  [-2, -1, 1, 2],
  [-2,  0, 0, 2],
  [-1,  0, 0, 1]
]
💡
Sort the array.
Fix one element arr[i] (like i=0 in the example).
Now the problem becomes: find all triplets in the remaining array that sum to target - arr[i].
Outer loop: for i = 0 .. n-4
Fix arr[i]
Compute a new target for the remaining array:
newTarget = target - arr[i]
Call 3Sum logic on arr[i+1:] with newTarget

3Sum logic: use the usual two-pointer approach to find triplets whose sum = newTarget.

Combine arr[i] with each triplet to form a quadruplet.

*/

func main() {
	fmt.Println(Foursum([]int{1, 0, -1, 0, -2, 2},0))
}


func Foursum(arr []int,target int) [][]int {
	result := [][]int{}
	sort.Ints(arr)
	for i:=0;i<len(arr)-2;i++ {
		if i>0 && arr[i] == arr[i-1] {
			continue
		}
		Newtarget := target - arr[i]
		triplet := ThreeSum(arr[i+1:],Newtarget)
		for _, triplet := range triplet {
			result1 := append([]int{arr[i]},triplet...)
			result = append(result, result1)
		}
		

	}
	return result
}

func ThreeSum(arr []int, target int) [][]int {
	sort.Ints(arr)
	result := [][]int{}
	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == target {
				result = append(result, []int{arr[i], arr[left], arr[right]})
				left++
				right--
				for left < right && arr[left] == arr[left-1] {
					left++
				}
				for left < right && arr[right] == arr[right+1] {
					right--
				}

			} else if sum < target {
				left++
				for left < right && arr[left] == arr[left-1] {
					left++
				}
			} else {
				right--
				for left < right && arr[right] == arr[right+1] {
					right--
				}
			}

		}
	}
  return result
}
