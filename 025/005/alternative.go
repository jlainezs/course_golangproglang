package main

import (
	"fmt"
	"sort"
)

type userAlt struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := userAlt{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := userAlt{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := userAlt{
		First: "Jaime",
		Last:  "Bonet",
		Age:   32,
		Sayings: []string{
			"boooo",
			"more booo",
		},
	}

	users := []userAlt{u1, u2, u3}

	for _, u := range users {
		sort.Strings(u.Sayings)
	}

	// since go 1.8
	sort.Slice(users, func(i, j int) bool {
		if users[i].Age < users[j].Age {
			return true
		}

		return users[i].Last < users[j].Last
	})

	for _, u := range users {
		fmt.Printf("%s %s %v\n", u.First, u.Last, u.Age)
		for _, s := range u.Sayings {
			fmt.Printf("\t%s\n", s)
		}
	}
}
