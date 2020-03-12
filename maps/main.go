package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Hex value for %s is %s\n", color, hex)
	}
}