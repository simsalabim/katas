package main

import "strconv"

func problem_004() int {
	max_palindrome := 0

	for n1 := 999; n1 > 99; n1-- {
		for n2 := 999; n2 > 99; n2-- {
			if is_palindrome_number(n1*n2) && n1*n2 > max_palindrome {
				max_palindrome = n1 * n2
			}
		}
	}
	return max_palindrome
}

func is_palindrome_number(n int) bool {
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
