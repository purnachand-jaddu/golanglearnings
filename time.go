package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()
	timeAsString := fmt.Sprintf("%d_%d_%d_%d_%d_%d", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), time.Second())
	fmt.Println(timeAsString)
}
