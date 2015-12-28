package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	v := float64(rand.Intn(100))
	user_defined := false

	if len(os.Args) > 1 {
		input_v, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Println(os.Stderr, "An error has occurred", err)
			os.Exit(1)
		}
		v = input_v
		user_defined = true
	}

	if !user_defined {
		fmt.Println("Sqrt of", v)
	}
	fmt.Println(Sqrt(v))
	fmt.Println(math.Sqrt(v))
}

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}
