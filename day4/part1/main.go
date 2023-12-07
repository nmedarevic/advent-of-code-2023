package main

import (
	avent_math "advent_helper/advent_math"
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

func findIntersection(array1 *[]int, array2 *[]int) *[]int {
	var longerArray = array1
	var shorterArray = array2

	if len(*array2) > len(*array1) {
		longerArray = array2
		shorterArray = array1
	}

	var result []int

	for _, item := range *longerArray {
		for _, item2 := range *shorterArray {
			if item == item2 {
				result = append(result, item)
			}
		}
	}

	return &result
}

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
	commonNumbers = *findIntersection(winningNumbers, numbers)

	fmt.Println("Winning numbers:", winningNumbers)
	fmt.Println("Numbers:", numbers)
	fmt.Println("Common numbers:", commonNumbers)

	if len(commonNumbers) == 0 {
		return 0, ""
	}

	return avent_math.IntPow(2, len(commonNumbers)-1), ""
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

	for ; workerNumber < 10; workerNumber++ {
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

	fmt.Println("Result array:", lineNumbers)

	sum := 0

	for _, number := range lineNumbers {
		sum += number
	}

	fmt.Println("=====================")
	fmt.Println("Result is", sum)
	fmt.Println("=====================")

	if totalLinesProcessed != len(lines) {
		fmt.Println("Not all lines processed!")
	}
}

// 20667
