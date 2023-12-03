package main

import (
	"advent_helper/file_loader"
	"advent_helper/strings_helpers"
	"advent_helper/workers"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

type GameColors struct {
	red   int
	blue  int
	green int
}
type Game struct {
	colors GameColors
	gameId int
}

var gameLimits = GameColors{red: 12, green: 13, blue: 14}

var numberPattern = "\\d+"
var redPattern = "\\d+ red"
var bluePattern = "\\d+ blue"
var greenPattern = "\\d+ green"

var numberRegex = regexp.MustCompile(numberPattern)
var redRegex = regexp.MustCompile(redPattern)
var blueRegex = regexp.MustCompile(bluePattern)
var greenRegex = regexp.MustCompile(greenPattern)

func parseLine(line string) (int, string) {
	var gameAndCombos = strings.Split(line, ":")

	var allPlays = strings.Split(gameAndCombos[1], ";")

	biggestRed := 0
	biggestBlue := 0
	biggestGreen := 0

	for _, play := range allPlays {
		redMatches := redRegex.FindString(play)
		blueMatches := blueRegex.FindString(play)
		greenMatches := greenRegex.FindString(play)

		redNumberString := numberRegex.FindString(redMatches)
		blueNumberString := numberRegex.FindString(blueMatches)
		greenNumberString := numberRegex.FindString(greenMatches)

		var red = strings_helpers.StringToNumberDefaultToZero(redNumberString)
		var blue = strings_helpers.StringToNumberDefaultToZero(blueNumberString)
		var green = strings_helpers.StringToNumberDefaultToZero(greenNumberString)

		if red > biggestRed {
			biggestRed = red
		}

		if blue > biggestBlue {
			biggestBlue = blue
		}

		if green > biggestGreen {
			biggestGreen = green
		}
	}

	return biggestRed * biggestBlue * biggestGreen, ""
}

func getCallback(result *[]int) func(string) {
	return func(line string) {
		fmt.Println(line)
		gameId, error := parseLine(line)

		if error != "" {
			return
		}

		*result = append(*result, gameId)
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

// Result = 70768
