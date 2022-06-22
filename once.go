package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	target := 18
	iterations := 100
	wg := sync.WaitGroup{}
	once := sync.Once{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go checkIfWeHitTheTarget(&target, &wg, &once)
	}
	wg.Wait()

}

func checkIfWeHitTheTarget(value *int, wg *sync.WaitGroup, once *sync.Once) {
	hitValue := rand.Intn(20)
	checkIfWeHit(hitValue, value, once)
	wg.Done()
}

func checkIfWeHit(hitValue int, value *int, once *sync.Once) {
	if hitValue == *value {
		once.Do(func() {
			fmt.Println("Target Achieved")
		})
	}
}
