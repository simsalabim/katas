package main

var cache map[int]int = make(map[int]int)

func problem_002(n int) int {
	cache[1] = 1
	cache[2] = 1

	sum := 0
	i := 1
	for 1 == 1 {
		f := fib(i)
		if f > n {
			break
		}
		if f%2 == 0 {
			sum += f
		}
		i++
	}
	return sum
}

func fib(n int) int {
	if cached_value, ok := cache[n]; ok {
		return cached_value
	}
	cache[n] = fib(n-1) + fib(n-2)
	return cache[n]
}
