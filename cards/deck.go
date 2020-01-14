package main

import "fmt"

// deck now inherits all functionalities of a string slice
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
