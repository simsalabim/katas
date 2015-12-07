package main

func problem_006(limit int) int {
	sum_of_squares, sum := 0, 0

	for i := 1; i <= limit; i++ {
		sum_of_squares += i * i
		sum += i
	}

	return sum*sum - sum_of_squares
}
