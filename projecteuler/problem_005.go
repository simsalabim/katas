package main

func problem_005(dividers []int) int {
	for i := 1; ; i++ {
		if is_divisible(i, dividers) {
			return i
		}
	}
}

func is_divisible(n int, dividers []int) bool {
	for _, divider := range dividers {
		if n%divider != 0 {
			return false
		}
	}
	return true
}
