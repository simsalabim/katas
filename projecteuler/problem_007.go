package main

func problem_007(n int) int {
	counter := 0
	for i := 0; ; i++ {
		if is_prime(i) {
			counter++
		}
		if counter == n {
			return i
		}
	}
}
