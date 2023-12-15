package main

import (
	"advent_helper/file_loader"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var cardMap = map[string]int{
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
	bid      int
	strength int
	cards    string
	rank     int
}

func CalculateStrength(cards string) int {
	strength := 0

	var frequencyMap = map[string]int{}
	for _, card := range cards {
		frequencyMap[string(card)]++
	}

	for key, value := range frequencyMap {
		if value == 5 {
			return 6
		}

		if value == 4 {
			return 5
		}

		if value == 3 {
			hasTwoDifferentOnes := false

			for _, value1 := range frequencyMap {
				if value1 == value {
					continue
				}

				// Full house
				if value1 == 2 {
					return 4
				}

				// Three of a kind
				if value1 == 1 {
					hasTwoDifferentOnes = true
				}
			}

			// Three of a kind
			if hasTwoDifferentOnes {
				return 3
			}
		}

		// Two pair
		if value == 2 {
			hasAnotherPair := false
			for key1, value1 := range frequencyMap {
				if key == key1 {
					continue
				}

				if value1 == 2 {
					hasAnotherPair = true
				}
			}

			if hasAnotherPair {
				return 2
			}
		}

		// One pair
		if value == 2 {
			allAreOnes := false

			for key1, value1 := range frequencyMap {
				if key1 == key {
					continue
				}
				if value1 == 1 {
					allAreOnes = true
				}
				if value1 != 1 {
					allAreOnes = false
				}
			}

			if allAreOnes {
				return 1
			}
		}

		strength = 0
	}

	return strength
}

/*
*
1  - card 1 is stronger
0  - cards are equal
-1 - card 2 is stringer
*
*/
func compareTwoElements(card1 string, card2 string) int {
	for i := 0; i < 5; i++ {
		if cardMap[string(card1[i])] > cardMap[string(card2[i])] {
			return 1
		} else if cardMap[string(card1[i])] < cardMap[string(card2[i])] {
			return -1
		}
	}

	return 0
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

		bid, _ := strconv.Atoi(result[1])
		// fmt.Println(result, bid)

		hands = append(hands, Hand{bid: bid, cards: result[0], strength: CalculateStrength(result[0])})

		sizeOfHands := len(hands)

		if sizeOfHands == 1 {
			continue
		}

		if hands[sizeOfHands-2].strength > hands[sizeOfHands-1].strength {
			hands[sizeOfHands-2], hands[sizeOfHands-1] = hands[sizeOfHands-1], hands[sizeOfHands-2]
		}

		if hands[sizeOfHands-2].strength == hands[sizeOfHands-1].strength {
			comparison := compareTwoElements(hands[sizeOfHands-2].cards, hands[sizeOfHands-1].cards)

			if comparison == 1 {
				hands[sizeOfHands-2], hands[sizeOfHands-1] = hands[sizeOfHands-1], hands[sizeOfHands-2]
			}
		}
	}

	for i := 2; i <= len(hands); i++ {
		if hands[i-2].strength > hands[i-1].strength {
			hands[i-2], hands[i-1] = hands[i-1], hands[i-2]

		}

		if hands[i-2].strength == hands[i-1].strength {
			comparison := compareTwoElements(hands[i-2].cards, hands[i-1].cards)

			if comparison == 1 {
				hands[i-2], hands[i-1] = hands[i-1], hands[i-2]
			}
		}

		hands[i-2].rank = i - 1
		hands[i-1].rank = i
	}

	var result = 0

	for _, hand := range hands {
		fmt.Println(hand.cards, hand.rank)
		result += hand.bid * hand.rank
	}
	fmt.Println(result)
}
