package main

import (
	"advent_helper/file_loader"
	"advent_helper/strings_helpers"
	"fmt"
	"regexp"
)

var seedsRegexPattern = `seeds: (\d+[ ]?)+`
var filtersRegexPattern = `(\-?\w+\-?)+ \w+:\n((\d+ ?)+\n)+`
var onlyNumbersRegexPattern = `((\d+ ?)+\n)`
var numbersRegexpattern = `\d+`

var seedsRegex = regexp.MustCompile(seedsRegexPattern)
var filtersRegex = regexp.MustCompile(filtersRegexPattern)
var onlyNumbersRegex = regexp.MustCompile(onlyNumbersRegexPattern)
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

func main() {
	var wholeFile = file_loader.LoadFileAsString("./input_short.txt")

	var seedsNumbers = *getSeeds(&wholeFile)

	filtersMatches := filtersRegex.FindAllString(wholeFile, -1)

	var filters = make([][]Filter, len(filtersMatches))

	for index, matches := range filtersMatches {
		filters[index] = make([]Filter, 0)

		filters[index] = append(filters[index], *createFilter(matches)...)
	}

	// fmt.Println("ALL FILTERS", filters)

	// for _, seed := range seedsNumbers {

	// }

	var filtersLevelOne = *createFilter(filtersMatches[0])

	var seedsLevelOne = []int{}

	seedsLevelOne = *passThroughFilter(seedsNumbers, filtersLevelOne)

	fmt.Println(filtersLevelOne)
	fmt.Println(seedsLevelOne)
}
