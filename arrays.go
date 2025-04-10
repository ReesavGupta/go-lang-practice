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

	// maps in go

	mymap := map[string]uint64{"reesav": 2, "roshan": 5}

	fmt.Printf("\nthis is the map: %v", mymap)

	// we can get the map's value as:-
	lol, ok := mymap["pritam"]
	//if the key is not present then maps will return the "zero" value of it specific type

	// maps when we try to get the value, it returns an "ok" boolean variable which states if the variable is present or not
	if ok {
		fmt.Printf("\nthis is lol: %v", lol)
	} else {
		fmt.Printf("\nthe key is not present")
	}

	// to iterate over a map or an array you can use the for range or the traditional for

	for key, value := range mymap {
		fmt.Printf("\nThis is the key: %v and this is its value: %v", key, value)
	}

	// you can create a while loop in go as:
	i := 0
	for {
		if i >= 10 {
			fmt.Printf("\nwe are here")
			break
		}
		i++
		fmt.Printf("\nThis is i: %d", i)
	}
}
