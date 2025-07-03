package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("CPUs\t\t%d\n", runtime.NumCPU())
	fmt.Printf("Goroutines\t%d\n", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mx.Lock()
			v := counter
			v++
			counter = v
			mx.Unlock()
			wg.Done()
		}()
		fmt.Printf("Goroutines\t%d\n", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Printf("Goroutines\t%d\n", runtime.NumGoroutine())

	// shows 1, bc we are rewriting it
	fmt.Printf("Counter\t%d\n", counter)
}
