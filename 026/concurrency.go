package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("OS\t\t%s\n", runtime.GOOS)
	fmt.Printf("ARCH\t\t%s\n", runtime.GOARCH)

	fmt.Printf("CPUs\t\t%d\n", runtime.NumCPU())
	fmt.Printf("Goroutines\t\t%d\n", runtime.NumGoroutine())
	wg.Add(1)
	go f1()

	wg.Add(1)
	f2()

	fmt.Printf("CPUs\t\t%d\n", runtime.NumCPU())
	fmt.Printf("Goroutines\t\t%d\n", runtime.NumGoroutine())
	wg.Wait() // wait go routines to end
}

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("f1 - %d\n", i)
	}
	wg.Done()
}

func f2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("f2 - %d\n", i)
	}
	wg.Done()
}
