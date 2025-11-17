package main

import "fmt"

/*
Given two strings s1 and s2 that may contain the '#' character,
determine if they are equal after applying all backspaces.

The '#' means “backspace” → delete the previous character if one exists.
Input: s1 = "xy#z", s2 = "xzz#"
Output: true
Explanation:
s1 → "xz"
s2 → "xz"
Instead of actually building new strings,
we can traverse both from the end → start using two pointers:
i = len(s1)-1   → last index
j = len(s2)-1
Now, move both pointers from right to left:
If you see a #, increment a skip counter and move left.
If you see a letter but still have a skip > 0, skip that letter and move left.

Stop when you’re on a valid character (one that isn’t removed).

When both pointers are on valid characters:

Compare them → if different → return false.

Else move both left and continue.

If both finish → return true.
*/

func main() {
	fmt.Println(TwoPointerSolution("xy#z", "xzz#")) // true
	fmt.Println(TwoPointerSolution("ab#c", "ad#c")) // true
	fmt.Println(TwoPointerSolution("xy#z", "xyz#")) // false
}

// Calling this by passing the index one by one to get the valid index this will get the valid one without backspace and the one backspaced
func NextValidIndex(s string, currentIndex int) int {
	skip := 0
	for currentIndex >= 0 {
		if s[currentIndex] == '#' {
			skip++
			currentIndex--
		} else if skip > 0 {
			currentIndex--
			skip--
		} else {
			break
		}

	}
	return currentIndex
}

func TwoPointerSolution(s1, s2 string) bool {
	i, j := len(s1)-1, len(s2)-1

	for i >= 0 || j >= 0 {
		i = NextValidIndex(s1, i)
		j = NextValidIndex(s2, j)

		if i < 0 && j < 0 {
			return true
		}

		if i < 0 || j < 0 {
			return false
		}

		if s1[i] != s2[j] {
			return false
		}
		i--
		j--
	}
	return true
}

func build(s string) []rune {
	var result []rune // space complexity o(n)

	for _, ch := range s {
		if ch == '#' {
			result = result[:len(result)-1]
		} else {
			result = append(result, ch)
		}
	}
	return result
}

func SimpleSolution(s1, s2 string) bool {
	b1 := build(s1)
	b2 := build(s2)
	return string(b1) == string(b2)
}
