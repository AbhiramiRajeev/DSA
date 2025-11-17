/*
Given an array of integers, find the length of the smallest subarray which, if sorted, would make the whole array sorted.
arr := []int{1, 2, 5, 3, 7, 10, 9, 12}
✅ Output → 5
(Because sorting [5, 3, 7, 10, 9] makes the whole array sorted.)
Find first out-of-order element from the left
Move left forward until you find arr[left] > arr[left+1].

In example →
1, 2 ✅
5 > 3 ❌ → left = 2

Find first out-of-order element from the right
Move right backward until you find arr[right] < arr[right-1].

In example →
12, 9 ✅
10 < 9 ❌ → right = 6

Find min and max inside the window
window = arr[left:right+1]
windowMin = 3
windowMax = 10
Expand left boundary
Move left back while elements before it are greater than windowMin.

Expand right boundary
Move right forward while elements after it are smaller than windowMax.

Answer = right - left + 1
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 3, 7, 10, 9, 12}
	fmt.Println(MinWindowSort(arr))
}

func MinWindowSort(arr []int) int {
    n := len(arr)
    left, right := 0, n-1

    // Step 1: find first out-of-order from left
    for left < n-1 && arr[left] <= arr[left+1] {
        left++
    }
    if left == n-1 {
        return 0 // fully sorted
    }

    // Step 2: find first out-of-order from right
    for right > 0 && arr[right] >= arr[right-1] {
        right--
    }

    // Step 3: find min & max in window
    windowMin, windowMax := arr[left], arr[left]
    for i := left; i <= right; i++ {
        if arr[i] < windowMin {
            windowMin = arr[i]
        }
        if arr[i] > windowMax {
            windowMax = arr[i]
        }
    }

    // Step 4: expand left boundary
    for left > 0 && arr[left-1] > windowMin {
        left--
    }

    // Step 5: expand right boundary
    for right < n-1 && arr[right+1] < windowMax {
        right++
    }

    return right - left + 1
}
