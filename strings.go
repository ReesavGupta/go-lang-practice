package main

import (
	"fmt"
	"strings"
)

func stringss() {
	// fmt.Printf("this is inside strings")
	var mystr string = "hello"

	for i, v := range mystr {
		fmt.Printf("\nthis is the index: %d and this is the value: %v", i, v)
	}

	var strBuilder strings.Builder

	myStrSlice := []string{"h", "e", "l", "l", "o"}

	for i := range myStrSlice {
		strBuilder.WriteString(myStrSlice[i])
	}

	catStr := strBuilder.String()

	fmt.Printf("\nthis is the cat string: %s", catStr)
}
