package main

import (
	"fmt"
	"sync"
)

var elems = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(len(elems))
	// wg.Add(len(elems) + 1) // Infinite wait time
	// wg.Add(len(elems) - 1) // Negative wait group counter
	for _, elem := range elems {
		go process(elem, &wg)
	}
	wg.Wait()
}

func process(elem int, wg *sync.WaitGroup) {
	fmt.Printf("Processing %d", elem)
	fmt.Println()
	wg.Done()
}
