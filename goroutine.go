package main

import (
	"sync"
)

func getProductConcurrent() []string {
	r := make([]string, 0, 100)

	semaphore := make(chan struct{}, 10)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for id := 0; id < 50; id++ {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(id int) {
			defer func() {
				wg.Done()
				<-semaphore
			}()

			res, ok := searchProduct(id)

			mu.Lock()
			defer mu.Unlock()
			if ok {
				r = append(r, res)
			}
		}(id)
	}
	wg.Wait()
	return r
}
