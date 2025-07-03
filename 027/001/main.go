package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sub1() {
	defer wg.Done()
	fmt.Printf("Sub 1\n")
}

func sub2() {
	defer wg.Done()
	fmt.Printf("Sub 2\n")
}

func main() {
	wg.Add(2)
	go sub1()
	go sub2()
	wg.Wait()
	fmt.Printf("main finished\n")
}
