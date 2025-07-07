package main

import (
	"fmt"
	"sync"
)

func createRoutines(c chan string) {
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < 10; j++ {
					c <- fmt.Sprintf("%d - %d", i, j)
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(c)
	}()
}

func main() {
	c := make(chan string)
	createRoutines(c)

	for v := range c {
		fmt.Printf("%s\n", v)
	}

	fmt.Println("Done!")
}
