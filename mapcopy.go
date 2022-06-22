package main

import "fmt"

func main() {
	fMap := make(map[int]bool)
	fMap[0] = true
	fMap[1] = false
	sMap := fMap
	sMap[2] = true
	fmt.Println(&fMap)
	fmt.Println(&sMap)

	fSlice := []int{1,2}
	sSlice := fSlice
	fmt.Println(&fSlice[1])
	fmt.Println(&sSlice[1])
	sSlice[1] = 3
	fmt.Println(&sSlice[1])
	fmt.Println(sSlice[1])
}