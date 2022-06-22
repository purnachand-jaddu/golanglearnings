package main

import (
	"fmt"
	"math"
)

type Type1 bool
type Type2 bool
type Type3 Type1

type I interface {
	SayHI()
}

type C struct {
}

func (c C) SayHI() {
	fmt.Println("Hello")
}

func saySomeThing(i I) {
	i.SayHI()
}

func main() {
	saySomeThing(C{})
}

func someFunc() {
	var first Type1 = false
	//var third Type3 = first -> This will not work
	var third Type3 = false
	fmt.Println(first)
	fmt.Println(third)
}

func findUnsortedSubarray(nums []int) int {
	indicesStack := []int{}
	leftMost := math.MaxInt32
	rightMost := math.MinInt
	for i := 0; i < len(nums); i++ {
		if len(indicesStack) == 0 {
			indicesStack = append(indicesStack, i)
		} else {
			topIndex := indicesStack[len(indicesStack)-1]
			if nums[i] >= nums[topIndex] {
				indicesStack = append(indicesStack, i)
			} else {
				rightMost = max(rightMost, i)
				popped := -1
				for {
					if len(indicesStack) == 0 {
						break
					}
					topI = indicesStack[len(indicesStack)-1]
					if nums[i] < nums[topI] {
						popped := indicesStack[len(indicesStack)-1]
						indicesStack = indicesStack[:len(indicesStack)-1]
					} else {
						break
					}
				}
			}
		}
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
