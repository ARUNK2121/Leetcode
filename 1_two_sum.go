package main

// 1 : Two Sum
// Time Complexity : O(n)
// Space Complexity : O(n)

func twoSum(nums []int, target int) []int {
	c := make(map[int]int)
	for i, v := range nums {
		if j, ok := c[target-v]; ok {
			if i != j {
				return []int{j, i}
			}
		}
		c[v] = i
	}
	return []int{}
}
