package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for range 50 {
		wg.Go(func() {
			for range 1000 {
				ops.Add(1)
			}
		})
	}

	wg.Wait()
	println("ops:", ops.Load())
}
