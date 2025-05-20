package main

import "fmt"

type ByteSize uint64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB

	// Overflows
	//ZB
	//YB
)

func main() {
	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t\t %b\n", TB, TB)
	fmt.Printf("%d \t\t\t %b\n", PB, PB)
	fmt.Printf("%d \t\t\t %b\n", EB, EB)
}
