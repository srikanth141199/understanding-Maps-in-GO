package main

import "fmt"

func main() {
	color := map[string]string{
		"red":   "#ff0000",
		"green": "4bf745",
	}

	fmt.Println(color)
	//diff ways to create a map
	var color1 map[string]string
	color2 := make(map[string]string)
	fmt.Println(color1)
	fmt.Println(color2)
}
