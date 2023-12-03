package main

import (
	"advent_helper/file_loader"
	"fmt"
	"regexp"
)

var escapedDot = regexp.QuoteMeta(".")
var escapedNumerical = regexp.QuoteMeta("d")
var specialCharacterPattern = "[^" + escapedDot + "\\d+]"
var specialCharacterRegex = regexp.MustCompile(specialCharacterPattern)

func main() {
	var lines = file_loader.LoadLinesFromFile("../part1/input.txt")
	var specialChars = make(map[string]string)

	for _, line := range lines {
		var allCharsInLine = specialCharacterRegex.FindAllString(line, -1)

		for _, char := range allCharsInLine {
			_, ok := specialChars[char]

			if !ok {
				specialChars[char] = char
			}
		}
	}

	for key, _ := range specialChars {
		fmt.Println(key)
	}
}

//* % @ - / + =

/**
"#", "=", "/", "&", "%", "@", "$", "-", "*"
**/
