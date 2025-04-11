package main

import (
	"fmt"
	"time"
)

// channels hold data.
// channels are thread safe.
// we can listen for data when added or removed

var c = make(chan int) // unbuffered channel
// var c = make(chan int, 5)  // c : [0,1,2,3,4]
// buffered channel can hold multiple values without popping out the current value.

func channels() {

	/*
		while running the code like this we get :

		// fatal error: all goroutines are asleep - deadlock!

		fmt.Printf("\n")
		c := make(chan int) // unbuffered channel

		c <- 5 // c : [5]

		--> when we write into an UB channel the code blocks until something else reads from it

		--> thus we will be waiting on line no. 16 forever. unable to reach the line 24 where we actually read from the channel

		--> luckily go's runtime is smart enough to notice this and not freeze the execution but throw a deadlock

		i := <-c // c : [] , i = 5

		fmt.Printf("this is the value of channel : %v", i)

	*/
	go process()

	// i := <-c
	for i := range c {
		fmt.Printf("\nthis is i : %v", i)
		time.Sleep(time.Second * 1) // does some work which takes a second
	}

}

func process() {
	defer close(c) // we will need to close the channel before the func ends as the loop in the above channels function will keep on waiting for a response. and will remain stuck forever again.
	for i := range 5 {
		c <- i
	}
	fmt.Printf("\nExiting the process ...")

	// when using a buffered channel: the process() is called from the main func.
	// all the values ranged over 5 are added to the buffered channel. and the process func prints "exiting the preocess" and exits.
	// now what happens in the main func is that since there are 5 values in the buffered channel the we start to read the values off the channel with a second of delay sugggesting that we are doing some work which takes a second

	// but when we use a unbuffered channel what happens is we add a value to the buffer and we read the value out of it concurrently. which makes the rest of the process function wait.

}
