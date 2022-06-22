package main

import (
	"fmt"
	"time"
)

func SumUsingTwoChannels() {
	start := time.Now()
	input := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	length := len(input)
	go noteSumBetweenIndices(0, length/2, &input, c)
	go noteSumBetweenIndices((length/2)+1, length-1, &input, c)
	time.Sleep(time.Second)
	x, y := <-c, <-c
	fmt.Printf("sum is %d", x+y)
	timeElapsed := time.Since(start)
	fmt.Printf("Total time taken %d", timeElapsed)
}

func noteSumBetweenIndices(start int, end int, sliceAddr *[]int, c chan int) {
	sum := 0
	slice := *sliceAddr
	for i := start; i <= end; i++ {
		sum += slice[i]
	}
	c <- sum
}
