package main

import "fmt"

func Clear() {
	for i := 0; i < 30; i++ {
		fmt.Println()
	}
}

func PrintWord() {
	var c CurrentWord
	if current == nil {
		return
	}
	c = *current

	// print word
	Clear()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("   ", "- - - - - - - - - ")
	fmt.Println("   ", c.Original)
	fmt.Println("   ", c.Meta.Revealed)
	fmt.Println("   ", "- - - - - - - - - ")
	fmt.Println("   ", c.Similar)
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
