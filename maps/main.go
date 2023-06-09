package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#00000",
		"white": "#fffff",
	}

	// fmt.Println(colors)
	printMap(colors)

}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, hex)
	}
}
