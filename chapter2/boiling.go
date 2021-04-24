package main

import "fmt"

const freezeF, boilingF = 32.0, 212.0

func main() {
	fmt.Printf("boiling point = %g F or %g C\n", boilingF, fToC(boilingF))
	fmt.Printf("freezing point = %g F or %g C\n", freezeF, fToC(freezeF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
