package main

import "strconv"

// Fuzz Buzz
// Time Complexity : O(n)
// Space Complexity : O(n)

func fizzBuzz(n int) []string {
	r := make([]string, 0)
	i := 1
	for i <= n {
		r = append(r, GetStringForNumber(i))
		i++
	}
	return r
}

func GetStringForNumber(n int) string {
	if n%3 == 0 {
		if n%15 == 0 {
			return "FizzBuzz"
		} else {
			return "Fizz"
		}
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
