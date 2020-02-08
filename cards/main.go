package main

import "fmt"

func main() {

	cards := newDeckFromFile("my_cards.txt")
	cards.shuffle()

	fmt.Println(cards)
}
