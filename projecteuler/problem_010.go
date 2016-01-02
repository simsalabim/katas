package main

func problem_010(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if is_prime(i) {
			sum += i
		}
	}
	return sum
}
