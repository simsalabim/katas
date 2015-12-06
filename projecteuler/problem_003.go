package main

import "math"

func problem_003(n int) int {
	for i := int(math.Sqrt(float64(n))); i > 1; i-- {
		if n%i == 0 && is_prime(i) {
			return i
		}
	}
	return 1
}

func is_prime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
