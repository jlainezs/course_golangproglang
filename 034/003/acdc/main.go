// Package acdc, starting Another highway to hell
package acdc

// Sum sums all the given integers.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
