package main

import (
	"advent_helper/file_loader"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

var numbersRegexpattern = `\d+`

var numberRegex = regexp.MustCompile(numbersRegexpattern)

type Filter struct {
	input  []uint64
	output []uint64
}

type SeedRange struct {
	start uint64
	end   uint64
}

func main() {
	readFile := file_loader.OpenFile("./input.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	fileScanner.Scan()
	line := fileScanner.Text()

	var inputSeeds []uint64 = []uint64{}

	matches := numberRegex.FindAllString(line, -1)

	for _, m := range matches {
		v, _ := strconv.ParseUint(m, 10, 64)
		inputSeeds = append(inputSeeds, v)
	}

	var seeds []SeedRange = []SeedRange{}
	for i := 0; i < len(inputSeeds); i += 2 {

		seeds = append(seeds, SeedRange{start: inputSeeds[i], end: inputSeeds[i+1]})
	}

	fmt.Println("Seeds done:", seeds)

	var filters = make([][]Filter, 7)

	fileScanner.Scan()
	for i := 0; i < 7; i++ {
		fileScanner.Scan()

		filters[i] = make([]Filter, 0)

		for fileScanner.Scan() {
			line = fileScanner.Text()
			if line == "" {
				break
			}

			matches = numberRegex.FindAllString(line, -1)
			output, _ := strconv.ParseUint(matches[0], 10, 64)
			input, _ := strconv.ParseUint(matches[1], 10, 64)
			filterLength, _ := strconv.ParseUint(matches[2], 10, 64)

			var lineFilter = Filter{
				input: []uint64{
					input,
					input + filterLength,
				},
				output: []uint64{
					output,
					output + filterLength,
				},
			}

			filters[i] = append(filters[i], lineFilter)
		}
	}

	fmt.Println("Filters done:", filters)

	var i uint64 = 0
	var minimalValue uint64 = 0
	var originalNumber uint64 = 0

completeLoop:
	for ; ; i++ {
		originalNumber = i
		minimalValue = i

		var filterClassIndex int = len(filters) - 1
		// fmt.Println("NEW START")
		// Go from classes in reverse
	withNextFilterClass:
		for ; filterClassIndex >= 0; filterClassIndex-- {
			// Go from filters in reverse
			var filterIndex = len(filters[filterClassIndex]) - 1
			for ; filterIndex >= 0; filterIndex-- {
				if minimalValue >= filters[filterClassIndex][filterIndex].output[0] &&
					minimalValue <= filters[filterClassIndex][filterIndex].output[1] {
					// fmt.Println("Here", minimalValue, filters[filterClassIndex][filterIndex].input[0]+(minimalValue-filters[filterClassIndex][filterIndex].input[0]))
					minimalValue = filters[filterClassIndex][filterIndex].input[0] + (minimalValue - filters[filterClassIndex][filterIndex].output[0])
					continue withNextFilterClass
				}
			}
		}

		for _, seedRange := range seeds {
			if minimalValue >= seedRange.start && minimalValue <= seedRange.start+seedRange.end {
				// fmt.Println("AAAAA", minimalValue, seedRange)
				break completeLoop
			}
		}
	}

	fmt.Println(minimalValue, originalNumber)

	// 	for _, filterClass := range filters {
	// 		for _, filter := range filterClass {
	// 			if seeds[index] >= filter.input[0] && seeds[index] <= filter.input[1] {
	// 				// fmt.Println("BEFORE", seeds[index], filter.input[0], filter.output[0])
	// 				seeds[index] = filter.output[0] + (seeds[index] - filter.input[0])
	// 				// fmt.Println("AFTER", seeds[index])
	// 				continue withNextFilterClass
	// 			}
	// 		}
	// 	}
	// }

	// for index := range seeds {
	// withNextFilterClass:
	// 	for _, filterClass := range filters {
	// 		for _, filter := range filterClass {
	// 			if seeds[index] >= filter.input[0] && seeds[index] <= filter.input[1] {
	// 				// fmt.Println("BEFORE", seeds[index], filter.input[0], filter.output[0])
	// 				seeds[index] = filter.output[0] + (seeds[index] - filter.input[0])
	// 				// fmt.Println("AFTER", seeds[index])
	// 				continue withNextFilterClass
	// 			}
	// 		}
	// 	}
	// 	// fmt.Println("seeds[index]")
	// }

	// // fmt.Println("seeds", seeds)

	// fmt.Println("Lowest number:", array_helpers.FindMin[uint64](seeds))
}

// 6082852
