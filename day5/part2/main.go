package main

import (
	"advent_helper/array_helpers"
	"advent_helper/file_loader"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

var numbersRegexpattern = `\d+`

var numberRegex = regexp.MustCompile(numbersRegexpattern)

type Filter struct {
	input       []uint64
	output      []uint64
	filterIndex int
}

func main() {
	readFile := file_loader.OpenFile("./input.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	fileScanner.Scan()
	line := fileScanner.Text()

	var seeds []uint64 = []uint64{}

	matches := numberRegex.FindAllString(line, -1)

	for _, m := range matches {
		v, _ := strconv.ParseUint(m, 10, 64)
		seeds = append(seeds, v)
	}

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
				filterIndex: i,
			}

			filters[i] = append(filters[i], lineFilter)
		}
	}

	for index := range seeds {

	withNextFilterClass:
		for _, filterClass := range filters {
			for _, filter := range filterClass {

				if seeds[index] >= filter.input[0] && seeds[index] <= filter.input[1] {
					// fmt.Println("BEFORE", seeds[index], filter.input[0], filter.output[0])
					seeds[index] = filter.output[0] + (seeds[index] - filter.input[0])
					// fmt.Println("AFTER", seeds[index])
					continue withNextFilterClass
				}
			}
		}
		// fmt.Println("seeds[index]")
	}

	// fmt.Println("Seeds", seeds)

	fmt.Println("Lowest number:", array_helpers.FindMin[uint64](seeds))
}
