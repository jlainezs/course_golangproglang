package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Printf("CPUs\t\t%d\n", runtime.NumCPU())
	fmt.Printf("Goroutines\t%d\n", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Printf("Counter %d\n", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Printf("Goroutines\t%d\n", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Printf("Goroutines\t%d\n", runtime.NumGoroutine())

	// shows 1, bc we are rewriting it
	fmt.Printf("Counter\t%d\n", counter)
}
