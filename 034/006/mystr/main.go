// Package mystr does things over strings
package mystr

import "strings"

// Cat returns the concatenation of all given strings, separated with a white space
func Cat(xs []string) string {
	s := ""
	for _, v := range xs {
		s += v
		s += " "
	}

	return s
}

// Join uses strings.Join over the given strings using " " as separator
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
