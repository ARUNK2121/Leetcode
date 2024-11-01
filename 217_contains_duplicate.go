package main

// Contains Duplicate
// Time Complexity : O(n)
// Space Complexity : O(n)

func containsDuplicate(nums []int) bool {
	c := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := c[v]; ok {
			return true
		}
		c[v] = struct{}{}
	}

	return false
}
