# Pointers

As in C:

- `&` to get the address of a variable
- `*` means pointer

```go
package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x) // value
	fmt.Println(&x) // address
	fmt.Printf("%T\n", &x) // it says *int, meaning is a pointer to an integer
}
```

NOTE: int is not and *int, int refers to a an integer value while *int refers to the memory address of an integer value.

## Dereferencing pointers

```go
package main

import "fmt"

func main() {
	x := 42
	y := &x
	fmt.Printf("%T\n", &y) // it says *int
	fmt.Println(*y) // it says the value, 42
	
	*y = 43
	fmt.Println(*y) // it says the value, 43
	fmt.Println(x) // it also says the value 43, as y is the address of x.
}
```

## Pass by value / reference

In go, built in types are passed by value (copy) except the reference types:

1. Pointers
2. Slices
3. Maps
4. Channels
5. Functions
6. Interfaces

By "reference type" we mean some structure thats passed by reference.
