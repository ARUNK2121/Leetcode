package main

// Valid Anagram
// Time Complexity : O(n)
// Space Complexity : O(n)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	z := map[rune]int{}
	for _, v := range s {
		z[v]++
	}
	for _, v := range t {
		z[v]--
		if z[v] < 0 {
			return false
		}
	}

	return true
}
