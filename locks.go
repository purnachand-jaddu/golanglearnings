package main

import (
	"fmt"
	"sync"
)

func main() {
	resource := 0
	iterations := 1000
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go incrementResource(&resource, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(resource)
}

func incrementResource(v *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	lock.Lock()
	*v = *v + 1
	lock.Unlock()
	wg.Done()

}
