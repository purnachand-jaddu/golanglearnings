package main

import (
	"fmt"
	"time"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	channel := make(chan int)
	go iterateOverValues(values, channel, 0)
	for {
		value, ok := <-channel
		if !ok {
			break
		}
		fmt.Println(value)
	}
}

func iterateOverValues(values []int, channel chan int, index int) {
	if index >= len(values) {
		close(channel)
		return
	}
	channel <- values[index]
	iterateOverValues(values, channel, index+1)
}

func undestandingChannelsWithZeroCapacity() {
	channel := make(chan bool)
	go createResponseFile1(channel)
	fmt.Println(<-channel)
}

func createResponseFile1(channel chan bool) {
	channel <- true
	time.Sleep(time.Second)
}
