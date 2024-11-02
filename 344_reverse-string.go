package main

// Reverse String
// Time Complexity : O(n)
// Space Complexity : O(1)

func reverseString(s []byte) {
	i := 0
	l := len(s) - 1
	for i <= (l / 2) {
		s[i], s[l-i] = s[l-i], s[i]
		i++
	}
}
