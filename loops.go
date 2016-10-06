package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
    t, z, c := 0., 1., 0
	for {
	    c++
		z, t = z - (math.Pow(z, 2) - x) / (2 * z), z
		if math.Abs(t-z) < 1e-8 {
            break
		}
	}
	return z, c
}

func main() {
	value, count := Sqrt(2)
    fmt.Println(count)
    fmt.Println(value == math.Sqrt(2))
}
