package main

import (
	"fmt"
)

func main() {
	color := map[string]string{
		"red":   "#ff0000",
		"green": "4bf745",
		"white": "#fffff",
	}

	fmt.Println(color)
	printMap(color)
	//diff ways to create a map
	var color1 map[string]string
	color2 := make(map[string]string)
	fmt.Println(color1)
	fmt.Println(color2)
}

// we can iterate in map since it has indexes
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
