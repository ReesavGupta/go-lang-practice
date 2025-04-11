package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}
var wg = sync.WaitGroup{}
var m = sync.Mutex{}

func goRoutiens() []string {
	t0 := time.Now()
	for i := range dbData {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()

	fmt.Printf("\nTotal execution time : %v", time.Since(t0))
	return results
}

func dbCall(i int) {
	// simulate db call
	var delay float32 = 2000

	time.Sleep(time.Duration(delay) * time.Millisecond)

	// fmt.Printf("\nthe result from the data is : %v", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}
