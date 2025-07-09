package main

import "fmt"

func Average(numbers []float64) float64 {
	result := 0.0
	for _, v := range numbers {
		result += v
	}

	return result / float64(len(numbers))
}

func main() {
	r := Average([]float64{1, 2})
	fmt.Println(r)
}
