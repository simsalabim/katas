package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const precision = 1e-10

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
	fmt.Printf("Newton with precesion %e: %g\n", precision, Sqrt(v))
	fmt.Println("Stdlib:", math.Sqrt(v))
}

func Sqrt(x float64) float64 {
	z := x / 2
	prev_z := z
	for {
		z = z - (z*z-x)/(2*z)
		if prev_z-z <= precision {
			break
		}
		prev_z = z
	}
	return z
}
