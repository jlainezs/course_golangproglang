package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	// to defer a function which returns an error
	// we defer an anonymous function which does the
	// error management
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("error %s", err)
		}
	}(f) // the function will be called with this arguments

	s := []byte("Hello World")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
