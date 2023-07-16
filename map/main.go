package main

import "fmt"

func main() {
	// other ways to declare map
	// 1. var colors map[string]string
	/* 2. colors := make(map[int]string)
	colors[10] = "#ffffff" */

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
	
	/* delete(colors, 10)
	fmt.Println(colors)*/
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
