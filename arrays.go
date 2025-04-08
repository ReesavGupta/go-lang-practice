package main

import "fmt"

func arrays() {
	// arrays are of fixed size whereaas the sizes of a slice is variable
	var arrSlice = []int64{1, 2, 3}

	var myArr = [...]string{"hello", "what are you doing"}

	// myArr = append(myArr, "what is up nigguh") doesn't work

	fmt.Printf("this is arrSlice before append: %v\n", arrSlice)

	arr := append(arrSlice, 5) // works like a charm

	fmt.Printf("\nthis is arrSlice: %v", arr)

	fmt.Printf("\nthis is string myArr: %v", myArr)
}
