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
var starRegex = regexp.MustCompile(`\*`)

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
			horizontal: match[1],
		},
	})

	fmt.Println("Added a number:", numberInt)
}

func getNumbersNearChars(lines []string, outputArray *[]Number) *[]Number {
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
					saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
					continue continueLoop
				}

				// Check above between two numbers
				for i := match[0]; i < match[1]; i++ {
					if isSpecialCharacter(lines[lineIndex-1][i]) {
						// fmt.Println(lines[lineIndex-1][i])
						saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
						continue continueLoop
					}
				}

				// Check top diagonal up
				// Only if the match is not on the last index
				if match[1] < len(lines[lineIndex-1]) && isSpecialCharacter(lines[lineIndex-1][match[1]]) {
					// fmt.Println(line[match[0]:match[1]], string(lines[lineIndex-1][match[1]]))
					saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
					continue continueLoop
				}
			}

			// Check left
			// If match is not on the first index
			if match[0] > 0 && isSpecialCharacter(line[match[0]-1]) {
				// fmt.Println(string(lines[lineIndex][match[0]-1]))
				saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
				continue continueLoop
			}

			// Check right
			// If match is not on the last index
			if match[1] < len(line) && isSpecialCharacter(line[match[1]]) {
				// fmt.Println(string(lines[lineIndex][match[1]+1]))
				saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
				continue continueLoop
			}

			// If not the last line
			if lineIndex+1 <= numberOfLines-1 {
				// Check left diagonal down
				// Only if the match does not start at the first index
				if match[0] > 0 && isSpecialCharacter(lines[lineIndex+1][match[0]-1]) {
					// fmt.Println(string(lines[lineIndex+1][match[0]-1]))
					saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
					continue continueLoop
				}

				// Check below
				// Check above between two numbers
				for i := match[0]; i < match[1]; i++ {
					if isSpecialCharacter(lines[lineIndex+1][i]) {
						// fmt.Println(string(lines[lineIndex+1][i]))
						saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
						continue continueLoop
					}
				}

				// Check right diagonal dowm
				if match[1] < len(lines[lineIndex+1]) && isSpecialCharacter(lines[lineIndex+1][match[1]]) {
					// fmt.Println(string(lines[lineIndex+1][match[0]+1]))
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
	if len(*outputArray) > 0 && len((*outputArray)[len(*outputArray)-1]) == 1 {
		var alreadyFoundNumber = (*outputArray)[len(*outputArray)-1][0]

		if alreadyFoundNumber.startPosition.vertical == vertical &&
			alreadyFoundNumber.startPosition.horizontal <= horizontal &&
			alreadyFoundNumber.endPosition.horizontal >= horizontal {
			return nil
		}
	}

	var number = getNumber(vertical, horizontal, inputArray)

	return number
}

func addNumberToOutputArray(number *Number, outputArray *[][]Number) {
	(*outputArray)[len(*outputArray)-1] = append((*outputArray)[len(*outputArray)-1], *number)
}

func findAllCogs(lines []string, numbersAdjecentToChars *[]Number, outputArray *[][]Number) *[][]Number {
	var numberOfLines = len(lines)

	for lineIndex, line := range lines {
		cogMatches := starRegex.FindAllStringSubmatchIndex(line, -1)

	continueLoop:
		for _, match := range cogMatches {
			*outputArray = append(*outputArray, []Number{})

			// If not the first line
			if lineIndex > 0 {
				// Check left diagonal up
				if match[0] > 0 {
					number := findNumber(lineIndex-1, match[0]-1, numbersAdjecentToChars, outputArray)
					// var alreadyFoundNumber = (*outputArray)[len(*outputArray)-1][0]

					// if alreadyFoundNumber.startPosition.vertical
					// var number = getNumber(lineIndex-1, match[0]-1, numbersAdjecentToChars)

					if number != nil {
						addNumberToOutputArray(number, outputArray)
						// (*outputArray)[len(*outputArray)-1] = append((*outputArray)[len(*outputArray)-1], *number)

						if len(*outputArray) == 2 {
							continue continueLoop
						}
					}
				}

				// Check above
				for i := match[0]; i < match[1]; i++ {
					number := findNumber(lineIndex-1, i, numbersAdjecentToChars, outputArray)

					if number != nil {
						addNumberToOutputArray(number, outputArray)

						if len(*outputArray) == 2 {
							continue continueLoop
						}
					}
				}

				// Check top diagonal up
				// Only if the match is not on the last index
				if match[1] < len(lines[lineIndex-1]) {
					number := findNumber(lineIndex-1, match[1], numbersAdjecentToChars, outputArray)

					if number != nil {
						addNumberToOutputArray(number, outputArray)

						if len(*outputArray) == 2 {
							continue continueLoop
						}
					}
				}
			}

			// Check left
			// If match is not on the first index
			if match[0] > 0 {
				number := findNumber(lineIndex, match[0]-1, numbersAdjecentToChars, outputArray)

				if number != nil {
					addNumberToOutputArray(number, outputArray)

					if len(*outputArray) == 2 {
						continue continueLoop
					}
				}
				// fmt.Println(string(lines[lineIndex][match[0]-1]))
				// saveNumber(line[match[0]:match[1]], lineIndex, match, outputArray)
			}

			// Check right
			// If match is not on the last index
			if match[1] < len(line) {
				number := findNumber(lineIndex, match[1], numbersAdjecentToChars, outputArray)

				if number != nil {
					addNumberToOutputArray(number, outputArray)

					if len(*outputArray) == 2 {
						continue continueLoop
					}
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

						if len(*outputArray) == 2 {
							continue continueLoop
						}
					}
				}

				// Check below
				// Check above between two numbers
				for i := match[0]; i < match[1]; i++ {
					number := findNumber(lineIndex+1, i, numbersAdjecentToChars, outputArray)

					if number != nil {
						addNumberToOutputArray(number, outputArray)

						if len(*outputArray) == 2 {
							continue continueLoop
						}
					}
				}

				// Check right diagonal dowm
				if match[1] < len(lines[lineIndex+1]) {
					number := findNumber(lineIndex+1, match[1], numbersAdjecentToChars, outputArray)

					if number != nil {
						addNumberToOutputArray(number, outputArray)

						if len(*outputArray) == 2 {
							continue continueLoop
						}
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

	for _, number := range numbersNearCogs {
		fmt.Print(number[0].number)

		if (len(number)) > 1 {
			fmt.Print(" * ")
			fmt.Print(number[1].number)
		}

		fmt.Println("=")
	}

	sum := 0
	for _, number := range numbersAdjecentToChars {
		sum += number.number
	}

	fmt.Println("Sum of all numbers adjecent to special characters:", sum)
}

// Attempt 1 -
