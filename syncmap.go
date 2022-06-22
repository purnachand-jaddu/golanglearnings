package main

import "sync"

func main() {
	regMap := make(map[int]interface{})
	syncMap := sync.Map{}
	lock := sync.Mutex{}
	iterations := 10
	wg := sync.WaitGroup{}
	wg.Add(2 * iterations)
	for i := 0; i < iterations; i++ {
		go storeInConcurrentMap(&syncMap, i, &wg)
		go storeInMap(regMap, i, &wg, &lock)
	}
	wg.Wait()
}

func storeInConcurrentMap(syncMap *sync.Map, index int, wg *sync.WaitGroup) {
	syncMap.Store(index, 1)
	wg.Done()
}

func storeInMap(regMap map[int]interface{}, index int, wg *sync.WaitGroup, lock *sync.Mutex) {
	lock.Lock()
	regMap[index] = 1
	lock.Unlock()
	wg.Done()
}
