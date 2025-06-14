package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func frequency(reader io.Reader) (map[string]int, error) {
	wordFreq := make(map[string]int)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		for _, w := range words {
			wordFreq[w]++
		}
	}

	return wordFreq, nil
}

/*
Count words
*/
func main() {
	f, err := os.Open("moby_dick.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer f.Close()

	freq, readError := frequency(f)
	if readError != nil {
		log.Fatalf("error: %s", readError)
	}

	for k, v := range freq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
