package main

import (
	"fmt"
	"main/dog"
)

func main() {
	dogYears := dog.Years(20)
	fmt.Printf("%d\n", dogYears)
}

// install godoc
//     go install golang.org/x/tools/cmd/godoc@latest
// run it
//     $ godoc -http=:8080
// open a browser to http://localhost:8080
