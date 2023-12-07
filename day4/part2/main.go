package main

import (
	"advent_helper/array_helpers"
	"advent_helper/file_loader"
	"advent_helper/strings_helpers"
	"advent_helper/workers"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

var numberPattern = "\\d+"

var numberRegex = regexp.MustCompile(numberPattern)

func convertStringArrayToNumberArray(stringArray []string) *[]int {
	var result []int

	for _, item := range stringArray {
		result = append(result, strings_helpers.StringToNumberAndPanic(item))
	}

	return &result
}

func parseLine(line string) (int, string) {
	var gameAndNumbers = strings.Split(line, ":")

	// fmt.Println("Split line:", gameAndNumbers)
	if len(gameAndNumbers) == 1 {
		fmt.Println(line)
		fmt.Println(gameAndNumbers)
	}
	var winningNumbersAndNumbers = strings.Split(gameAndNumbers[1], "|")

	var commonNumbers []int

	winningNumberStrings := numberRegex.FindAllString(winningNumbersAndNumbers[0], -1)
	numberStrings := numberRegex.FindAllString(winningNumbersAndNumbers[1], -1)

	var winningNumbers = convertStringArrayToNumberArray(winningNumberStrings)
	var numbers = convertStringArrayToNumberArray(numberStrings)
	commonNumbers = *array_helpers.FindIntersection(winningNumbers, numbers)

	// fmt.Println("Winning numbers:", winningNumbers, "\nNumbers:", numbers, "\nCommon numbers:", commonNumbers)

	return len(commonNumbers), ""
}

func getCallback(result *[]int) func(string) {
	return func(line string) {
		score, error := parseLine(line)

		if error != "" {
			return
		}

		*result = append(*result, score)
	}
}

func main() {
	var lines = file_loader.LoadLinesFromFile("./input.txt")

	wg := &sync.WaitGroup{}
	tasks := make(chan []string)
	results := make(chan []int)

	var workerNumber int = 0

	for ; workerNumber < 1; workerNumber++ {
		wg.Add(1)
		go workers.Worker(wg, workerNumber, getCallback, tasks, results)
	}

	var batchOffset = 0
	var maxOffset = 10

	if len(lines) < 10 {
		maxOffset = len(lines)
	}

	for batchOffset < len(lines) {
		endOffset := batchOffset + maxOffset

		if endOffset > len(lines) {
			endOffset = len(lines)
		}

		tasks <- lines[batchOffset:endOffset]
		batchOffset += maxOffset
	}

	close(tasks)

	go workers.MonitorWorker(wg, results)

	totalLinesProcessed := 0

	var lineNumbers []int

	for item := range results {
		totalLinesProcessed += len(item)
		lineNumbers = append(lineNumbers, item...)
	}

	cards := map[int]int{}

	fmt.Println("Result array:", lineNumbers)
	fmt.Println("=====================")

	sum := 0

	for index, number := range lineNumbers {
		cards[index]++
		for numberOfCopies := 0; numberOfCopies < cards[index]; numberOfCopies++ {
			for i := index + 1; i <= index+number; i++ {
				cards[i]++
			}
		}
	}

	for _, number := range cards {
		sum += number
	}

	if totalLinesProcessed != len(lines) {
		fmt.Println("Not all lines processed!")
	}

	if sum == 5833065 {
		fmt.Println("=====================")
		fmt.Println("End result:", sum)
		fmt.Println("=====================")
	} else {
		fmt.Println("Try again")
	}
}

// incorrect 5833023
// 5833065 - correct
