package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negativ number: %g", float64(e))
}


func Sqrt(x float64) (float64, int, error) {
	if x < 0 {
		return x, 0, ErrNegativeSqrt(x)
	}
    t, z, c := 0., 1., 0
	for {
	    c++
		z, t = z - (math.Pow(z, 2) - x) / (2 * z), z
		if math.Abs(t-z) < 1e-8 {
            break
		}
	}
	return z, c, nil
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
