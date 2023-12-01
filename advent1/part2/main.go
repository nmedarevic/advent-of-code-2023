package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
)

var regexPattern = "\\d|one|two|three|four|five|six|seven|eight|nine|ten"

func loadInput() []string {
	var filePath = "./input.txt"

	readFile, err := os.Open(filePath)

	defer readFile.Close()

	if err != nil {
		fmt.Println(err)

		return []string{}
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}

func getLineNumber(numberMatches []string) int {
	if len(numberMatches) == 0 {
		return 0
	}

	if len(numberMatches) == 1 {
		number, err := strconv.Atoi(numberMatches[0])

		if err != nil {
			number = 0
		}

		return (number * 10) + number
	}

	if len(numberMatches) > 1 {
		firstNumber, err := strconv.Atoi(numberMatches[0])

		if err != nil {
			firstNumber = 0
		}

		secondNumber, err := strconv.Atoi(numberMatches[len(numberMatches)-1])

		if err != nil {
			secondNumber = 0
		}

		return (firstNumber * 10) + secondNumber
	}

	return 0
}

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func convertMatchesToNumbers(matches []string) []string {
	var result []string

	for _, match := range matches {
		value, ok := numberMap[match]

		if ok {
			result = append(result, value)
		} else {
			result = append(result, match)
		}
	}

	return result
}

func elementExistsInStringArray(array []string, element string) bool {
	for _, arrayElement := range array {
		if arrayElement == element {
			return true
		}
	}

	return false
}

func findAllNumbers(line string, r *regexp.Regexp) []string {
	var matches []string

	for i := 0; i < len(line); i++ {
		var substring = line[i:len(line)]
		match := r.FindString(substring)

		// if elementExistsInStringArray(matches, match) {
		// 	continue
		// }

		if line == "nine6kfpkqhkjzsknrldfcghcgkghnine" {
			fmt.Println(match, substring)
		}

		if match == "" {
			continue
		}

		matches = append(matches, match)
	}

	return matches
}

func getCallback(result *[]int) func(string) {
	return func(line string) {
		r := regexp.MustCompile(regexPattern)

		matches := findAllNumbers(line, r)

		matchedNumbers := convertMatchesToNumbers(matches)

		number := getLineNumber(matchedNumbers)

		*result = append(*result, number)

		if line == "nine6kfpkqhkjzsknrldfcghcgkghnine" {

			fmt.Println("Line:", line, "\n", "Matches:", matches, "\n", "Matched numbers:", matchedNumbers, "\n", "Final number:", number)
		}
	}
}

func iterate(lines <-chan []string, cb func(string)) {
	for task := range lines {
		for _, line := range task {
			cb(line)
		}
	}
}

func worker(wg *sync.WaitGroup, id int, lines <-chan []string, resultChanel chan<- []int) {
	defer wg.Done()

	fmt.Println("Worker", id, "started!")

	var result []int

	var callback = getCallback(&result)
	iterate(lines, callback)

	fmt.Println("Worker", id, "finished!")

	resultChanel <- result
}

func monitorWorker(wg *sync.WaitGroup, cs chan []int) {
	wg.Wait()
	close(cs)
}

func main() {
	var lines = loadInput()

	wg := &sync.WaitGroup{}
	tasks := make(chan []string)
	results := make(chan []int)

	var workerNumber int = 0

	for ; workerNumber < 10; workerNumber++ {
		wg.Add(1)
		go worker(wg, workerNumber, tasks, results)
	}

	var batchOffset = 0

	for batchOffset < len(lines) {
		tasks <- lines[batchOffset : batchOffset+10]
		batchOffset += 10
	}

	close(tasks)

	go monitorWorker(wg, results)

	totalLinesProcessed := 0

	var lineNumbers []int

	for item := range results {
		totalLinesProcessed += len(item)
		lineNumbers = append(lineNumbers, item...)

	}

	fmt.Println("Result:", lineNumbers)

	sum := 0

	for _, number := range lineNumbers {
		sum += number
	}

	fmt.Println("Result is", sum)

	if totalLinesProcessed != len(lines) {
		fmt.Println("Not all lines processed!")
	}
}

// 54522 - incorrect
// 54538 - incorrect
// 48360 - incorrect
// 53316 - incorrect

// Mozda? 54530
