package main

func problem_009(sum int) (int, int, int, int) {
	var a float32

	for b := 1; b < sum/2; b++ {
		a = float32((sum*sum/2 - sum*b) / (sum - b))
		if is_pythagorian(int(a), b, sum-int(a)-b) {
			return int(a), b, sum - int(a) - b, int(a) * b * (sum - int(a) - b)
		}
	}
	return -1, -1, -1, -1
}

func is_pythagorian(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}
