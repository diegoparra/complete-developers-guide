package main

import (
	"fmt"
)

type colorSet struct {
	Color string
	Hex   string
}

func main() {
	colorsMap := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf263",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// fmt.Println(colors)

	colorsMap["black"] = "#fa0000"

	printMap(colorsMap)

	colorsStruct := colorSet{
		Color: "red",
		Hex:   "#ff0000",
	}

	printStruct(colorsStruct)

}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("Color:", k)
		fmt.Println("HexValue:", v)
	}
}

func printStruct(s colorSet) {
	fmt.Println(s.Color)
	fmt.Println(s.Hex)
}
