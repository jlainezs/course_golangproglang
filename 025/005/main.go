package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type userSorting []user

func (s userSorting) Len() int      { return len(s) }
func (s userSorting) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s userSorting) Less(i, j int) bool {
	if s[i].Age < s[j].Age {
		return true
	}

	return s[i].Last < s[j].Last
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "Jaime",
		Last:  "Bonet",
		Age:   32,
		Sayings: []string{
			"boooo",
			"more booo",
		},
	}

	var users userSorting
	users = append(users, u1)
	users = append(users, u2)
	users = append(users, u3)

	for _, u := range users {
		sort.Strings(u.Sayings)
	}

	sort.Sort(users)

	for _, u := range users {
		fmt.Printf("%s %s %v\n", u.First, u.Last, u.Age)
		for _, s := range u.Sayings {
			fmt.Printf("\t%s\n", s)
		}
	}
}
