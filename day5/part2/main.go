package main

import (
	"advent_helper/array_helpers"
	"advent_helper/file_loader"
	"advent_helper/strings_helpers"
	"advent_helper/workers"
	"fmt"
	"regexp"
	"sync"
)

var seedsRegexPattern = `seeds: (\d+[ ]?)+`
var filtersRegexPattern = `(\-?\w+\-?)+ \w+:\n((\d+ ?)+\n)+`
var numbersRegexpattern = `\d+`

var seedsRegex = regexp.MustCompile(seedsRegexPattern)
var filtersRegex = regexp.MustCompile(filtersRegexPattern)
var numberRegex = regexp.MustCompile(numbersRegexpattern)

type Filter struct {
	input  []int
	output []int
}

func passThroughFilter(inputSeeds []int, filters []Filter) *[]int {
	var filteredSeeds = []int{}

withAnotherSeed:
	for _, seed := range inputSeeds {
		for _, filter := range filters {
			if seed >= filter.input[0] && seed <= filter.input[1] {
				filteredSeeds = append(filteredSeeds, filter.output[0]+(seed-filter.input[0]))

				continue withAnotherSeed
			}
		}
		filteredSeeds = append(filteredSeeds, seed)
	}

	return &filteredSeeds
}

func createFilter(filterInputString string) *[]Filter {
	filtersNumbers := numberRegex.FindAllString(filterInputString, -1)

	var filterArray = []Filter{}

	for i := 0; i < len(filtersNumbers); i += 3 {
		var filterLength = strings_helpers.StringToNumberAndPanic(filtersNumbers[i+2])
		var lineFilter = Filter{
			input: []int{
				strings_helpers.StringToNumberAndPanic(filtersNumbers[i+1]),
				strings_helpers.StringToNumberAndPanic(filtersNumbers[i+1]) + filterLength,
			},
			output: []int{
				strings_helpers.StringToNumberAndPanic(filtersNumbers[i]),
				strings_helpers.StringToNumberAndPanic(filtersNumbers[i]) + filterLength,
			},
		}

		filterArray = append(filterArray, lineFilter)
	}

	return &filterArray
}

func getSeeds(fileContents *string) *[]int {
	matchesSeeds := seedsRegex.FindAllString(*fileContents, -1)

	seeds := numberRegex.FindAllString(matchesSeeds[0], -1)

	var seedsNumbers = []int{}

	for _, seedString := range seeds {
		seedsNumbers = append(seedsNumbers, strings_helpers.StringToNumberAndPanic(seedString))
	}

	return &seedsNumbers
}

func createSeedArray(fileContents *string) *[]int {
	var seedsNumbersInitial = *getSeeds(fileContents)

	var seedsNumbers = []int{}

	for i := 0; i < len(seedsNumbersInitial); i += 2 {
		var seedStart = seedsNumbersInitial[i]
		var seedEnd = seedsNumbersInitial[i+1]

		for i := 0; i < seedEnd; i++ {
			seedsNumbers = append(seedsNumbers, seedStart+i)
		}
	}

	return &seedsNumbers
}

func createFilters(fileContents *string) *[][]Filter {
	filtersMatches := filtersRegex.FindAllString(*fileContents, -1)

	var filters = make([][]Filter, len(filtersMatches))

	for index, matches := range filtersMatches {
		filters[index] = make([]Filter, 0)

		filters[index] = append(filters[index], *createFilter(matches)...)
	}

	return &filters
}

func main() {
	var wholeFile = file_loader.LoadFileAsString("./input_short.txt")

	var seedsNumbers = *createSeedArray(&wholeFile)

	var filters = *createFilters(&wholeFile)

	var outputSeeds = make([]int, 0)

	outputSeeds = append(outputSeeds, seedsNumbers...)

	wg := &sync.WaitGroup{}
	tasks := make(chan []int)
	results := make(chan []int)

	var workerNumber int = 0
	var batch = 2

	for ; workerNumber < 10; workerNumber++ {
		wg.Add(1)

		go func(group *sync.WaitGroup, id int, intInput <-chan []int, resultChanel chan<- []int) {
			defer group.Done()
			fmt.Println("Worker", id, "started!")

			var resultArray []int

			for task := range intInput {

				var taskResult = make([]int, 0)
				taskResult = append(taskResult, task...)

				for _, filter := range filters {
					taskResult = *passThroughFilter(taskResult, filter)
				}

				resultArray = append(resultArray, taskResult...)
			}

			fmt.Println("Worker", id, "finished!", resultArray)

			resultChanel <- resultArray
		}(wg, workerNumber, tasks, results)
	}

	for i := 0; i < len(outputSeeds); i += batch {
		batchOffset := i + batch
		if batchOffset > len(outputSeeds) {
			batchOffset = len(outputSeeds)
		}

		tasks <- outputSeeds[i:batchOffset]
	}

	close(tasks)

	go workers.MonitorWorker(wg, results)

	var outputSeeds2 []int

	for outputArray := range results {
		outputSeeds2 = append(outputSeeds2, outputArray...)
	}

	fmt.Println("Lowest number:", array_helpers.FindLowestNumber((outputSeeds2)))

}
