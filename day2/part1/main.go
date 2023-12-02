package main

import (
	"advent_helper/file_loader"
	"advent_helper/workers"
	"fmt"
	"sync"
)

func getCallback(result *[]int) func(string) {
	return func(line string) {
		fmt.Println("Something")
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

	for batchOffset < len(lines) {
		tasks <- lines[batchOffset : batchOffset+10]
		batchOffset += 10
	}

	close(tasks)

	go workers.MonitorWorker(wg, results)

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
