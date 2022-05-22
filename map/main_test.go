package main

import "testing"

func BenchmarkMap(b *testing.B) {

	colorsMap := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf263",
	}

	printMap(colorsMap)

}

func BenchmarkStruct(b *testing.B) {

	colorsStruct := colorSet{
		Color: "red",
		Hex:   "#ff0000",
	}

	printStruct(colorsStruct)

}
