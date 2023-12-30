package main

import (
	"advent_helper/file_loader"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

var numberPattern = "\\d+"
var numberRegex = regexp.MustCompile(numberPattern)

func main() {
	filePath := "./input_short.txt"
	readFile := file_loader.OpenFile(filePath)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	collectiveResult := 0

	for {
		fileScanner.Scan()
		line := fileScanner.Text()

		if line == "" {
			break
		}

		inputArray := numberRegex.FindAllString(line, -1)
		inputArrayNumbers := make([]int, 0)

		for _, item := range inputArray {
			itemNumber, err := strconv.Atoi(item)

			if err != nil {
				panic(err)
			}

			inputArrayNumbers = append(inputArrayNumbers, itemNumber)
		}

		var matrix = make([][]int, 0)
		matrix = append(matrix, inputArrayNumbers)

		for {
			// fmt.Println(matrix[len(matrix)-1], len(matrix), matrix)
			outputArray := calculateResultArray(&matrix[len(matrix)-1], &matrix)
			matrix = append(matrix, *outputArray)
			// fmt.Println(matrix[len(matrix)-1], len(matrix), matrix)

			if arrayHasAllZeoros(outputArray) {
				break
			}
		}

		endItemResult := calculateEndRowResult(&matrix)
		collectiveResult += endItemResult

		fmt.Println(endItemResult)
	}

	fmt.Println(collectiveResult)
}

func calculateEndRowResult(matrix *[][]int) int {
	endItemResult := 0
	endItemArray := make([]int, 0)
	endItemArray = append(endItemArray, 0)

	for rowIndex := len(*matrix) - 2; rowIndex >= 0; rowIndex-- {
		// fmt.Println(matrix[rowIndex])
		endItemArray = append(endItemArray, endItemArray[len(endItemArray)-1]+(*matrix)[rowIndex][len((*matrix)[rowIndex])-1])
		endItemResult += endItemArray[len(endItemArray)-1]
	}

	return endItemArray[len(endItemArray)-1]
}

func arrayHasAllZeoros(input *[]int) bool {
	for _, item := range *input {
		if item != 0 {
			return false
		}
	}

	return true
}

func calculateResultArray(input *[]int, matrix *[][]int) *[]int {
	// (*matrix) = append((*matrix), []int{})
	var output = make([]int, 0)

	// currentLine := len(*matrix) - 1

	var difference = 0

	for i := 0; i < len(*input); i++ {
		if (i) == len(*input)-1 {
			break
		}

		difference = (*input)[i+1] - (*input)[i]

		output = append(output, difference)
	}

	// fmt.Println((*matrix)[currentLine])

	return &output
}
