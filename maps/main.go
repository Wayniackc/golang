package main

import "fmt"

func main() {

	// var colors map[string]string
	// Above does the same as below
	// colors := make(map[string]string)
	//
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b745f",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
