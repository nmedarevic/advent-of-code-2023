package main

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

func main() {
	text := `...............776..............552........968..................589...26...........484..............958......186....546.........484.........`

	sent := regexp.MustCompile("\\d+")
	matches := sent.FindAllStringSubmatchIndex(text, -1)

	for _, match := range matches {
		fmt.Println("match: ", text[match[0]:match[1]])
		fmt.Println("match: ", match)
		fmt.Println("rune", utf8.RuneCountInString(text[match[0]:match[1]]))
		// fmt.Println("context: ", text[utf8.RuneCountInString(text[:match[0]])])
		// // fmt.Println("next_tok: ", text[utf8.RuneCountInString(text[:match[4]]):utf8.RuneCountInString(text[:match[5]])])
		// fmt.Println("start: ", utf8.RuneCountInString(text[:match[2]]))
		// fmt.Println("end: ", utf8.RuneCountInString(text[:match[4]]))
		fmt.Println("------")
	}
}
