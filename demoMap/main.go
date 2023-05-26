package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colours["yellow"] = "xxxyyyy"
	delete(colours, "yellow")

	printColours(colours)
}

// Print all colours and its hex code
func printColours(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code of", colour, "is", hex)
	}
}
