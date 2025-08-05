package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 2.1

func main() {
	banner("go", 6)

	s := "g˚©"

	banner(s, 6)
	fmt.Println("len:", len(s))
	fmt.Println(":", s[1])

	for i, c := range s {
		fmt.Printf("%c at %d\n", c, i)
	}

}

func banner(text string, width int) {
	// BUG: len is in bytes
	// padding := (width - len(text)) / 2
	padding := (width - utf8.RuneCountInString(text)) / 2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", width))

}
