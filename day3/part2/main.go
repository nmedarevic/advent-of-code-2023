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

var specialCharsArray = []string{"#", "=", "/", "&", "%", "@", "$", "-", "*", "+"}

var numbersAdjecentToChars = []Number{}
var numbersNearCogs = [][]Number{}

var numberRegex = regexp.MustCompile(`\d+`)
var starRegex = regexp.MustCompile(`\*{1}`)

func isSpecialCharacter(value byte) bool {
	valueString := string(value)

	for _, specialChar := range specialCharsArray {
		if specialChar == valueString {
			return true
		}
	}

	return false
}

func saveNumber(numberString string, lineIndex int, match []int, numberArray *[]Number) {
	numberInt, err := strconv.Atoi(numberString)

	if err != nil {
		panic("Cannot convert the number to int")
	}

	*numberArray = append(*numberArray, Number{
		number: numberInt,
		startPosition: Position{
			vertical:   lineIndex,
			horizontal: match[0],
		},
		endPosition: Position{
			vertical:   lineIndex,
			horizontal: match[1] - 1,
		},
	})
}

func getNumbersNearChars(lines []string, outputArray *[]Number) *[]Number {
	var numberOfLines = len(lines)

	for lineIndex, line := range lines {
		numberPositionMatches := numberRegex.FindAllStringSubmatchIndex(line, -1)

	continueLoop:
		for _, match := range numberPositionMatches {
			// If not the first line
			if lineIndex > 0 {
				// Check left diagonal up
				if match[0] > 0 && isSpecialCharacter(lines[lineIndex-1][match[0]-1]) {
					saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
					continue continueLoop
				}

				// Check above between two numbers
				for i := match[0]; i < match[1]; i++ {
					if isSpecialCharacter(lines[lineIndex-1][i]) {
						saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
						continue continueLoop
					}
				}

				// Check top diagonal up
				// Only if the match is not on the last index
				if match[1] < len(lines[lineIndex-1]) && isSpecialCharacter(lines[lineIndex-1][match[1]]) {
					saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
					continue continueLoop
				}
			}

			// Check left
			// If match is not on the first index
			if match[0] > 0 && isSpecialCharacter(line[match[0]-1]) {
				saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
				continue continueLoop
			}

			// Check right
			// If match is not on the last index
			if match[1] < len(line) && isSpecialCharacter(line[match[1]]) {
				saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
				continue continueLoop
			}

			// If not the last line
			if lineIndex+1 <= numberOfLines-1 {
				// Check left diagonal down
				// Only if the match does not start at the first index
				if match[0] > 0 && isSpecialCharacter(lines[lineIndex+1][match[0]-1]) {
					saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
					continue continueLoop
				}

				// Check below
				// Check above between two numbers
				for i := match[0]; i < match[1]; i++ {
					if isSpecialCharacter(lines[lineIndex+1][i]) {
						saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
						continue continueLoop
					}
				}

				// Check right diagonal dowm
				if match[1] < len(lines[lineIndex+1]) && isSpecialCharacter(lines[lineIndex+1][match[1]]) {
					saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
					continue continueLoop
				}
			}
		}
	}

	return outputArray
}

func getNumber(vertical int, horizontal int, numbers *[]Number) *Number {
	for _, number := range *numbers {
		// If numbers are not in the same line
		if number.startPosition.vertical != vertical {
			continue
		}

		if number.startPosition.horizontal <= horizontal && number.endPosition.horizontal >= horizontal {
			return &number
		}
	}

	return nil
}

func findNumber(vertical int, horizontal int, inputArray *[]Number, outputArray *[][]Number) *Number {
	var number = getNumber(vertical, horizontal, inputArray)

	return number
}

func addNumberToOutputArray(number *Number, outputArray *[][]Number) {
	// Do not add existing numbers
	for _, existingNumber := range (*outputArray)[len(*outputArray)-1] {
		if existingNumber.startPosition.vertical == number.startPosition.vertical &&
			existingNumber.startPosition.horizontal == number.startPosition.horizontal &&
			existingNumber.endPosition.vertical == number.endPosition.vertical &&
			existingNumber.endPosition.horizontal == number.endPosition.horizontal {
			return
		}
	}

	(*outputArray)[len(*outputArray)-1] = append((*outputArray)[len(*outputArray)-1], *number)
	fmt.Print("Number ", number.number, "  ::::  ")
}

func findAllCogs(lines []string, numbersAdjecentToChars *[]Number, outputArray *[][]Number) *[][]Number {
	var numberOfLines = len(lines)

	for lineIndex, line := range lines {
		cogMatches := starRegex.FindAllStringSubmatchIndex(line, -1)

		// continueLoop:
		for _, match := range cogMatches {
			*outputArray = append(*outputArray, []Number{})
			fmt.Println("")
			fmt.Println("Coords", lineIndex, match)
			// If not the first line
			if lineIndex > 0 {
				// Check left diagonal up
				if match[0] > 0 {
					number := findNumber(lineIndex-1, match[0]-1, numbersAdjecentToChars, outputArray)

					if number != nil {
						addNumberToOutputArray(number, outputArray)
					}
				}

				// Check above
				number := findNumber(lineIndex-1, match[0], numbersAdjecentToChars, outputArray)

				if number != nil {
					addNumberToOutputArray(number, outputArray)
				}

				// Check right diagonal up
				// Only if the match is not on the last index
				if match[1] < len(lines[lineIndex-1]) {
					number := findNumber(lineIndex-1, match[0]+1, numbersAdjecentToChars, outputArray)

					if number != nil {
						addNumberToOutputArray(number, outputArray)
					}
				}
			}

			// Check left
			// If match is not on the first index
			if match[0] > 0 {
				number := findNumber(lineIndex, match[0]-1, numbersAdjecentToChars, outputArray)

				if number != nil {
					addNumberToOutputArray(number, outputArray)
				}
			}

			// Check right
			// If match is not on the last index
			if match[0]+1 < len(line) {
				number := findNumber(lineIndex, match[1], numbersAdjecentToChars, outputArray)

				if number != nil {
					addNumberToOutputArray(number, outputArray)
				}
			}

			// If not the last line
			if lineIndex+1 <= numberOfLines-1 {
				// Check left diagonal down
				// Only if the match does not start at the first index
				if match[0] > 0 {
					number := findNumber(lineIndex+1, match[0]-1, numbersAdjecentToChars, outputArray)

					if number != nil {
						addNumberToOutputArray(number, outputArray)
					}
				}

				// Check below
				number := findNumber(lineIndex+1, match[0], numbersAdjecentToChars, outputArray)

				if number != nil {
					addNumberToOutputArray(number, outputArray)
				}

				// Check right diagonal dowm
				if match[1] < len(lines[lineIndex+1]) {
					number := findNumber(lineIndex+1, match[1], numbersAdjecentToChars, outputArray)

					if number != nil {
						addNumberToOutputArray(number, outputArray)
					}
				}
			}
		}
	}

	return outputArray
}

func main() {
	var lines = file_loader.LoadLinesFromFile("./input.txt")

	numbersAdjecentToChars = *getNumbersNearChars(lines, &numbersAdjecentToChars)

	numbersNearCogs = *findAllCogs(lines, &numbersAdjecentToChars, &numbersNearCogs)

	fmt.Println("-------------------------")

	fmt.Println("Number of numbers:", len(numbersAdjecentToChars))
	fmt.Println("Numbers around cogs:", len(numbersNearCogs))

	sum := 0
	for _, number := range numbersNearCogs {
		if (len(number)) == 2 {
			sum += number[0].number * number[1].number
		}
	}

	if sum > 82818007 {
		fmt.Println("Sum of all cog numbers products:", sum)
	} else {
		fmt.Println("Try again")
	}
}

// Attempt 1 - 80788281 x

// 81035751 x - // Attempt 2 - to low

// 82818007 - CORRECT
