package main

import "fmt"

func main() {
	states := make([]string, 0, 50)
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))
	states = append(states, "Alabama", "Alaska", "Arizona", "Arkansas", "California",
		"Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho",
		"Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland",
		"Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska",
		"Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina",
		"North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina",
		"South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington",
		"West Virginia", "Wisconsin", "Wyoming")
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Printf("%v %s\n", i, states[i])
	}
}
