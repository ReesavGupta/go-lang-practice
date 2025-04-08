package main

import (
	"errors"
	"fmt"
)

func conditionals() (int, error) {
	var num1 int = 1
	var num2 int = 0

	fmt.Printf("this is num1 : %d and num2: %d \n", num1, num2)

	var err error

	var ans int

	if num2 == 0 {
		err = errors.New(`denominator is 0`)
		return 0, err
	} else {
		ans = num1 % num2
	}
	return ans, err
}
