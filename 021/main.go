package main

import "fmt"

type dog struct {
	first string
}

// a value receiver can only receive a value
func (d dog) walk() { // value receiver
	fmt.Println("My name is", d.first, "and I'm walking.")
}

// a pointer receiver can receive a pointer or a value of the type
func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

type athleticDog interface {
	walk()
	run()
}

func behaveLikeAnAthleticDog(d athleticDog) {
	d.run()
	d.walk()
}

func main() {
	d1 := dog{"Henry"}
	fmt.Printf("d1 type is %T\n", d1)
	d1.walk()
	d1.run()
	d1.walk()
	// The following statement raises a compiler error
	// 		cannot use d1 (variable of struct type dog) as athleticDog value in argument to behaveLikeAnAthleticDog:
	//		dog does not implement athleticDog (method run has pointer receiver)
	// Interface methods require to be called on pointers, not values
	// behaveLikeAnAthleticDog(d1)

	d2 := &dog{"Padget"}
	fmt.Printf("d2 type is %T\n", d2)
	d2.walk()
	d2.run()
	d1.walk()
	// here it works well, as
	behaveLikeAnAthleticDog(d2)
}
