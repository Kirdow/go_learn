package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	Section("Goroutines", runGoroutines)
}

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var diffs = []uint{}

func runGoroutines() {
	timeStart := time.Now()
	for range 1000 {
		wg.Add(1)
		go count(MyRandInt())
	}
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(timeStart))
	fmt.Printf("Hashes: %016x\n", diffs)
}

func count(start uint) {
	var res uint = start
	var valid bool
	for range 100000 {
		res, valid = MyDifficultHash(res, 24)
		if valid {
			m.Lock()
			diffs = append(diffs, res)
			m.Unlock()
		}
	}
	wg.Done()
}
