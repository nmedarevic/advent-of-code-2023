package main

import (
	"advent_helper/file_loader"
	"fmt"
	"regexp"
	"strconv"
)

type Position struct {
	vertical   int
	horizontal int
}
type Number struct {
	number        int
	startPosition Position
	endPosition   Position
}

var specialCharsArray = []string{"#", "=", "/", "&", "%", "@", "$", "-", "*"}

var numbersAdjecentToChars = []Number{}

var numberRegex = regexp.MustCompile("\\d+")

func isSpecialCharacter(value byte) bool {
	valueString := string(value)

	for _, specialChar := range specialCharsArray {
		if specialChar == valueString {
			return true
		}
	}

	return false
}

func saveNumber(numberString string, lineIndex int, match []int) {
	numberInt, err := strconv.Atoi(numberString)

	if err != nil {
		panic("Cannot convert the number to int")
	}

	numbersAdjecentToChars = append(numbersAdjecentToChars, Number{
		number: numberInt,
		startPosition: Position{
			vertical:   lineIndex,
			horizontal: match[0],
		},
		endPosition: Position{
			vertical:   lineIndex,
			horizontal: match[1],
		},
	})

	fmt.Println("Added a number:", numberInt)
}

func main() {
	var lines = file_loader.LoadLinesFromFile("./input.txt")

	var numberOfLines = len(lines)

	for lineIndex, line := range lines {
		numberPositionMatches := numberRegex.FindAllStringSubmatchIndex(line, -1)
		fmt.Println("-------------------------")
		fmt.Println("Line", lineIndex)

	continueLoop:
		for _, match := range numberPositionMatches {
			// If not the first line
			if lineIndex > 0 {
				// Check left diagonal up
				if match[0] > 0 && isSpecialCharacter(lines[lineIndex-1][match[0]-1]) {
					// fmt.Println(string(lines[lineIndex][match[0]-1]))
					saveNumber(line[match[0]:match[1]], lineIndex, match)
					continue continueLoop
				}

				// Check above between two numbers
				for i := match[0]; i < match[1]; i++ {
					if isSpecialCharacter(lines[lineIndex-1][i]) {
						// fmt.Println(lines[lineIndex-1][i])
						saveNumber(line[match[0]:match[1]], lineIndex, match)
						continue continueLoop
					}
				}

				// Check top diagonal up
				// Only if the match is not on the last index
				if match[1] < len(lines[lineIndex-1]) && isSpecialCharacter(lines[lineIndex-1][match[1]]) {
					// fmt.Println(line[match[0]:match[1]], string(lines[lineIndex-1][match[1]]))
					saveNumber(line[match[0]:match[1]], lineIndex, match)
					continue continueLoop
				}
			}

			// Check left
			// If match is not on the first index
			if match[0] > 0 && isSpecialCharacter(line[match[0]-1]) {
				// fmt.Println(string(lines[lineIndex][match[0]-1]))
				saveNumber(line[match[0]:match[1]], lineIndex, match)
				continue continueLoop
			}

			// Check right
			// If match is not on the last index
			if string(line[match[0]:match[1]]) == "873" {
				fmt.Println("873", match[0], match[1])
			}
			if match[1] < len(line) && isSpecialCharacter(line[match[1]]) {
				// fmt.Println(string(lines[lineIndex][match[1]+1]))
				saveNumber(line[match[0]:match[1]], lineIndex, match)
				continue continueLoop
			}

			// If not the last line
			if lineIndex+1 <= numberOfLines-1 {
				// Check left diagonal down
				// Only if the match does not start at the first index
				if match[0] > 0 && isSpecialCharacter(lines[lineIndex+1][match[0]-1]) {
					// fmt.Println(string(lines[lineIndex+1][match[0]-1]))
					saveNumber(line[match[0]:match[1]], lineIndex, match)
					continue continueLoop
				}

				// Check below
				// Check above between two numbers
				for i := match[0]; i < match[1]; i++ {
					if isSpecialCharacter(lines[lineIndex+1][i]) {
						// fmt.Println(string(lines[lineIndex+1][i]))
						saveNumber(line[match[0]:match[1]], lineIndex, match)
						continue continueLoop
					}
				}

				// Check right diagonal dowm
				if match[1] < len(lines[lineIndex+1]) && isSpecialCharacter(lines[lineIndex+1][match[1]]) {
					// fmt.Println(string(lines[lineIndex+1][match[0]+1]))
					saveNumber(line[match[0]:match[1]], lineIndex, match)
					continue continueLoop
				}
			}
		}
	}

	fmt.Println("-------------------------")

	fmt.Println("Number of numbers:", len(numbersAdjecentToChars))

	sum := 0
	for _, number := range numbersAdjecentToChars {
		sum += number.number
	}

	fmt.Println("Sum of all numbers adjecent to special characters:", sum)
}

// Attempt 2 = 474428 not correct
// Attempt 3 - 520249
