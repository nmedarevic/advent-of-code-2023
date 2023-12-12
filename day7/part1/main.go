package main

import (
	"advent_helper/file_loader"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var numbersRegexpattern = `\d+`
var cardPattern = `[AKQJT98765432]{5}`

var numberRegex = regexp.MustCompile(numbersRegexpattern)
var cardRegex = regexp.MustCompile(cardPattern)

var cardStrength = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

var carMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type Hand struct {
	bid   int
	cards string
}

func main() {
	readFile := file_loader.OpenFile("./input_short.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	var hands []Hand = []Hand{}
	for {
		fileScanner.Scan()
		line := fileScanner.Text()

		if line == "" {
			break
		}

		result := strings.Split(line, " ")
		fmt.Println(result)

		bid, _ := strconv.Atoi(result[1])

		hands = append(hands, Hand{bid: bid, cards: result[0]})
	}

	fmt.Println(hands)
}
