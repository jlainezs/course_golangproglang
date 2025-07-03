package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("ARCH: %s\n", runtime.GOARCH)
}
