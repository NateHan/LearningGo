package main

import "fmt"

func main() {
	// colors:= map[string]string {
	// 	"red": "#ff0000",
	// 	"green": "#4bf745"
	// }
	// fmt.Println(colors)

	// var colours map[string]string
	// fmt.Println(colours)

	// colors2 := make(map[string]string)
	// colors2["white"] = "#ffffff"
	// fmt.Println(colors2)
	// delete(colors2, "white")
	// fmt.Println(colors2)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Have the color %v, it is hex code %v. ", color, hex)
	}
}
