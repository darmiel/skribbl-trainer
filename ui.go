package main

import (
	"fmt"
)

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

	var str string
	var hr string
	var em string
	for _, c := range c.Meta.Revealed {
		if len(str) != 0 {
			str += " "
			hr += "═"
			em += " "
		}
		str += string(c)
		hr += "═"
		em += " "
	}

	// holy fuck,
	// this could be right from r/programminghorror
	// please fix this.
	var l1, l2, l3, l4, l5 string
	l1 = "1) " + c.Similar[0]
	i := len(c.Similar)
	if i >= 2 {
		l1 += ", 2) " + c.Similar[1]
	}

	if i >= 3 {
		l2 = "3) " + c.Similar[2]
	}
	if i >= 4 {
		l2 += ", 4) " + c.Similar[3]
	}

	if i >= 5 {
		l3 = "5) " + c.Similar[4]
	}
	if i >= 6 {
		l3 += ", 6) " + c.Similar[5]
	}

	if i >= 7 {
		l4 = "7) " + c.Similar[6]
	}
	if i >= 8 {
		l4 += ", 8) " + c.Similar[7]
	}

	if i >= 9 {
		l5 = "9) " + c.Similar[8]
	}
	if i >= 10 {
		l5 += ", 10) " + c.Similar[9]
	}

	fmt.Println("       ╔═"+hr+"═╗", l1)
	fmt.Println("       ║", em, "║", l2)
	fmt.Println("       ║", str, "║", l3)
	fmt.Println("       ║", em, "║", l4)
	fmt.Println("       ╚═"+hr+"═╝", l5)
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
