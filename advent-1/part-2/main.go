package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
)

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
	var resultNumber int = 0

	if len(numberMatches) == 0 {
		resultNumber = 0
	}

	if len(numberMatches) == 1 {
		number, err := strconv.Atoi(numberMatches[0])

		if err != nil {
			number = 0
		}

		resultNumber = (number * 10) + number
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

		resultNumber = (firstNumber * 10) + secondNumber
	}

	return resultNumber
}

func addRegexToCallback(r *regexp.Regexp, result *[]int) func(string) {
	return func(line string) {
		fmt.Println(line)
		matches := r.FindAllString(line, -1)
		fmt.Println(matches)

		number := getLineNumber(matches)

		fmt.Println("Found number for", matches, "is", number)

		*result = append(*result, number)
	}
}

func iterate(lines <-chan []string, cb func(string)) {
	for task := range lines {
		for _, line := range task {
			cb(line)
		}
	}
}

// func iterateThroughStringArrayChannel(lines <-chan []string, cb )

func worker(wg *sync.WaitGroup, id int, lines <-chan []string, resultChanel chan<- []int) {
	defer wg.Done()

	fmt.Println("Worker", id, "started!")

	r := regexp.MustCompile("\\d")

	var result []int

	// iterate(lines, )
	var callback = addRegexToCallback(r, &result)

	iterate(lines, callback)

	// for task := range lines {
	// 	for _, line := range task {
	// 		fmt.Println(line)
	// 		matches := r.FindAllString(line, -1)
	// 		fmt.Println(matches)

	// 		number := getLineNumber(matches)

	// 		fmt.Println("Found number for", matches, "is", number)

	// 		result = append(result, number)
	// 	}
	// }

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

		fmt.Println("Result:", item)
	}

	sum := 0

	for _, number := range lineNumbers {
		sum += number
	}

	fmt.Println("Result is", sum)

	if totalLinesProcessed != len(lines) {
		fmt.Println("Not all lines processed!")
	}
}
