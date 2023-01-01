package main

import (
	"fmt"
	"math"
)

type ErrNagativeSqrt float64

func (f ErrNagativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", f)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNagativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	val1, err1 := Sqrt(2)
	if err1 == nil {
		fmt.Printf("Sqrt is %f\n", val1)
	} else {
		fmt.Println(err1.Error())
	}

	val2, err2 := Sqrt((-2))
	if err2 == nil {
		fmt.Printf("Sqrt is %f\n", val2)
	} else {
		fmt.Println(err2.Error())
	}
}
