package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello World")
	fmt.Println(b.String())
	b.WriteString("Hello again.")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("The old content is gone.")
	fmt.Println(b.String())
}
