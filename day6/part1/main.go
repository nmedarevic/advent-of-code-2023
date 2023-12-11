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

func main() {
	readFile := file_loader.OpenFile("./input_short.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	fileScanner.Scan()
	lineTime := fileScanner.Text()
	fileScanner.Scan()
	lineDistance := fileScanner.Text()
	matchesTimes := numberRegex.FindAllString(lineTime, -1)
	matchesDistances := numberRegex.FindAllString(lineDistance, -1)

	var times []int = []int{}
	var distances []int = []int{}

	for _, item := range matchesTimes {
		number, _ := strconv.Atoi(item)
		times = append(times, number)
	}

	for _, item := range matchesDistances {
		number, _ := strconv.Atoi(item)
		distances = append(distances, number)
	}

	fmt.Println(times)
	fmt.Println(distances)

	var waysToWin []int = []int{0, 0, 0}

	for i := 0; i < 3; i++ {
		for time := 0; time < times[i]; time++ {
			var buttonTime = times[i] - time
			var distance = buttonTime * (times[i] - buttonTime)
			fmt.Println("button time", buttonTime, "distance", distance)

			if distance > distances[i] {
				waysToWin[i] += 1
			}
		}

		fmt.Println("--------")
	}

	fmt.Println("All distances multiplied", waysToWin[0]*waysToWin[1]*waysToWin[2])
}
