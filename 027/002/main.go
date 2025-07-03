package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Human interface {
	speak(whatToSay string)
}

func (p *Person) speak(whatToSay string) {
	fmt.Printf("%s says %s\n", p.name, whatToSay)
}

func saySomething(h Human) {
	h.speak("How do you do?")
}

func main() {
	p := Person{"Bob", 20}
	saySomething(&p) // this works
	//saySomething(p)	// this does not works
}
