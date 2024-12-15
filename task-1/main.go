package main

import (
	"errors"
	"fmt"
	"math"
)

func Calculator() (float64, error) {
	var a, b float64
	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	result := math.Sqrt(a*a + b*b)
	if result == 0 {
		return result, errors.New("triangle cant be 0")
	} else if result < 0 {
		return 0, errors.New("triangle cant be less than 0")
	}
	return result, nil

}
func main() {
	result, err := Calculator()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
