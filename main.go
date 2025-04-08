package main

import "fmt"

func fibo(num int) []int {
	var series []int = make([]int, num)

	for i := range num {

		if i == 0 {
			series[i] = 0
			continue
		} else if i == 1 {
			series[i] = 1
			continue
		}
		series[i] = series[i-1] + series[i-2]
	}
	return series
}

func main() {
	fmt.Println(`hello world`)

	var a float64 = 2
	var b float64 = 3

	fmt.Println(a / b)

	num := 10
	var fibSeries []int = fibo(num)

	fmt.Printf("This is the fibSeries for length %d: %v\n", num, fibSeries)
	result, err := conditionals()

	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		fmt.Printf("this is yout result %v", result)
	}

	fmt.Printf("\n*******this is arr*********")
	arrays()
}
