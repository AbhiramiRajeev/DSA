package main

import (
	"fmt"
	"math"
)

/*
Given an integer array arr and a target sum target,
find a triplet (three numbers) whose sum is closest to the target.

Return the sum of the triplet (not the triplet itself).
arr = [-2, 0, 1, 2], target = 2
Output:
1

Explanation:
Possible triplets and sums:

(-2, 0, 1) = -1

(-2, 0, 2) = 0

(-2, 1, 2) = 1

(0, 1, 2) = 3
Closest to 2 is 1.

Sort the array first.

For each number arr[i]:

Use two pointers:

left = i + 1

right = len(arr) - 1

Compute currentSum = arr[i] + arr[left] + arr[right]

Compare how close it is to the target:

If abs(target - currentSum) < abs(target - closestSum), update closestSum.

Move pointers:

If currentSum < target → move left++

Else → move right--

Return the closestSum after checking all possibilities.
*/
func main() {
	fmt.Println(TripletSumClose([]int{-2, 0, 1, 2}, 2))
}

func TripletSumClose(arr []int, target int) int {
	closestSum := arr[0] + arr[1] + arr[2]
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-closestSum)) {
				closestSum = sum
			}
			if sum < target {
				left++
			} else if sum == target {
				return sum
			} else {
				right--
			}
		}
	}
	return closestSum
}
