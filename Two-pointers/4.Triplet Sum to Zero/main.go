package main

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

func TripletSum()
