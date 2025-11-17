package main

import "fmt"

/*
Given an array containing 0s, 1s, and 2s, sort the array in-place so that all 0s come first, then all 1s, then all 2s.

You cannot use built-in sort — must solve it in O(n) time and O(1) space.
Input: arr = [2, 0, 2, 1, 1, 0]
Output: [0, 0, 1, 1, 2, 2]
💡 Logic (Three-Pointer / Two-Pointer Variation)

🔍 Step 1: Define three regions

We maintain three pointers:

low     → boundary for 0s
mid     → current element being examined
high    → boundary for 2s

Initially:

low = 0, mid = 0, high = len(arr)-1

And the array is conceptually divided like this:

[ 0s | 1s | unknown | 2s ]


[0s] → elements before low (sorted 0s)

[1s] → elements between low and mid-1 (sorted 1s)

[unknown] → elements between mid and high

[2s] → elements after high (sorted 2s)

🔍 Step 2: Traverse the array

While mid <= high:

If arr[mid] == 0:

Swap arr[low] and arr[mid]

low++, mid++

Why? We’re moving 0 to the 0-region, and mid moves forward.

If arr[mid] == 1:

Just mid++

Why? 1 is already in the correct middle region.

If arr[mid] == 2:

Swap arr[mid] and arr[high]

high-- (do NOT increment mid!)

Why? We move 2 to the 2-region. mid stays because the new element swapped from high needs to be examined.

🔍 Step 3: Stop

When mid > high, all elements are in correct positions.
*/

func main() {
	fmt.Println(DutchNationalFlag([]int{2, 0, 2, 1, 1, 0}))
}

func DutchNationalFlag(arr []int) []int {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] == 0 {
			arr[mid], arr[low] = arr[low], arr[mid]
			low++
			mid++ // why mid++ 
			/* After the swap:
			If arr[low] was 0 → swapping with 0 doesn’t change anything; moving mid++ is fine.
			If arr[low] was 1 → 1 is now at mid. That’s fine; mid++ moves past it, and it’s already in the correct middle region.
			If arr[low] was 2 → impossible before mid reaches it because all 2s are at the end (high */
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
	return arr
}
