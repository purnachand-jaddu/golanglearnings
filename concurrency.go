package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	names := []string{"A", "B", "C", "D"}
	channel := make(chan bool)
	for index, name := range names {
		go print(name, index, channel)
	}
	for i := 0; i < len(names); i++ {
		fmt.Println(<-channel)
	}
}

func print(s string, i int, c chan bool) {
	// time.Sleep(time.Second)
	fmt.Println(s)
	c <- true
}
