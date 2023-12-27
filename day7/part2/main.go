package main

import (
	"advent_helper/file_loader"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var cardMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

var StrengthMap = map[string]int{
	"FiveOfAKind":  6,
	"FourOfAKind":  5,
	"FullHouse":    4,
	"ThreeOfAKind": 3,
	"TwoPair":      2,
	"OnePair":      1,
	"HighCard":     0,
}

type Hand struct {
	bid      int
	strength int
	cards    string
}

type FrequencyItem struct {
	key   string
	value int
}

func partitionFrequency(arr *[]FrequencyItem, low, high int) (*[]FrequencyItem, int) {
	pivot := (*arr)[high]
	i := low
	for j := low; j < high; j++ {
		if (*arr)[j].value > pivot.value {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			i++
		}
	}

	(*arr)[i], (*arr)[high] = (*arr)[high], (*arr)[i]
	return arr, i
}

func quickSortHandsFrequecy(arr *[]FrequencyItem, low, high int) *[]FrequencyItem {
	if low < high {
		var p int
		arr, p = partitionFrequency(arr, low, high)
		arr = quickSortHandsFrequecy(arr, low, p-1)
		arr = quickSortHandsFrequecy(arr, p+1, high)
	}

	return arr
}

func CalculateStrength(cards string) int {
	// fmt.Println(cards)
	strength := StrengthMap["HighCard"]

	var frequencyMap = map[string]int{}

	for _, card := range cards {
		frequencyMap[string(card)]++
	}

	var frequencySortedArray = []FrequencyItem{}

	for key, value := range frequencyMap {
		frequencySortedArray = append(frequencySortedArray, FrequencyItem{key: key, value: value})
	}

	frequencySortedArray = *quickSortHandsFrequecy(&frequencySortedArray, 0, len(frequencySortedArray)-1)

	// highestFrequencyCard := "J"

	// for key, value := range frequencyMap {
	// 	if value >= cardMap[highestFrequencyCard] {
	// 		highestFrequencyCard = key
	// 	}
	// }

	// fmt.Println(highestFrequencyCard)
	// fmt.Println(frequencyMap)

	// var highestFrequencyCard = map[string]int{
	// 	frequencyMap["A"]: 0,
	// }

	// Add to all the number of Js
	// fmt.Println("BEFORE", frequencySortedArray)
	if frequencySortedArray[0].key != "J" {
		frequencySortedArray[0].value += frequencyMap["J"]

		for index, value := range frequencySortedArray {
			if value.key == "J" {
				frequencySortedArray[index].value = 0
			}
		}
	}
	// fmt.Println("AFTER", frequencySortedArray)

	// fmt.Println(frequencyMap)
	for _, item := range frequencySortedArray {
		// Five of a kind
		if item.value == 5 {
			strength = StrengthMap["FiveOfAKind"]

			return strength
		}

		// Four of a kind
		if item.value == 4 {
			if item.key != "J" {
				for keyInner := range frequencyMap {
					if keyInner == "J" {
						return StrengthMap["FiveOfAKind"]
					}
				}
			}

			return StrengthMap["FourOfAKind"]
			// strength = StrengthMap["FourOfAKind"]

			// return strength
		}

		// Full house
		// Three of a kind
		if item.value == 3 {
			hasTwoDifferentOnes := false

			for key1, value1 := range frequencyMap {
				if key1 == item.key {
					continue
				}

				// Full house
				if value1 == 2 {
					strength = StrengthMap["FullHouse"]

					if item.key != "J" {
						for keyInner := range frequencyMap {
							if keyInner == "J" {
								strength = StrengthMap["FourOfAKind"]
							}
						}
					}

					return strength
				}

				// Three of a kind
				if value1 == 1 {
					strength = StrengthMap["ThreeOfAKind"]
					hasTwoDifferentOnes = true
				}
			}

			// Three of a kind
			if hasTwoDifferentOnes {
				strength = StrengthMap["ThreeOfAKind"]

				if item.key != "J" {
					for keyInner := range frequencyMap {
						if keyInner == "J" {
							strength = StrengthMap["FullHouse"]
						}
					}
				}
			}

			return strength
		}

		// Two pair
		if item.value == 2 {
			hasAnotherPair := false
			for key1, value1 := range frequencyMap {
				if item.key == key1 {
					continue
				}

				if value1 == 2 {
					hasAnotherPair = true
				}
			}

			if hasAnotherPair {
				strength = StrengthMap["TwoPair"]

				if item.key != "J" {
					for keyInner := range frequencyMap {
						if keyInner == "J" {
							strength = StrengthMap["ThreeOfAKind"]
						}
					}
				}
			}
		}

		// One pair
		if item.value == 2 {
			allAreOnes := false

			for key1, value1 := range frequencyMap {
				if key1 == item.key {
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
				strength = StrengthMap["OnePair"]

				if item.key != "J" {
					for keyInner := range frequencyMap {
						if keyInner == "J" {
							strength = StrengthMap["TwoPair"]
						}
					}
				}
			}
		}
	}

	fmt.Println("++++++++", strength)
	return strength
}

/*
*
1  - card 1 is stronger
0  - cards are equal
-1 - card 2 is stringer
*
*/
func CompareTwoElements(card1 string, card2 string) int {
	for i := 0; i < 5; i++ {
		if cardMap[string(card1[i])] > cardMap[string(card2[i])] {
			return 1
		} else if cardMap[string(card1[i])] < cardMap[string(card2[i])] {
			return -1
		}
	}

	return 0
}

func partition(arr *[]Hand, low, high int) (*[]Hand, int) {
	pivot := (*arr)[high]
	i := low
	for j := low; j < high; j++ {
		if (*arr)[j].strength < pivot.strength || (*arr)[j].strength == pivot.strength && CompareTwoElements((*arr)[j].cards, pivot.cards) == -1 {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			i++
		}
	}

	// if (*arr)[i].cards == "T55J5" || (*arr)[high].cards == "T55J5" || (*arr)[i].cards == "QQQJA" || (*arr)[high].cards == "QQQJA" {
	// 	fmt.Println("REVERESE", (*arr)[i], (*arr)[high])
	// }
	(*arr)[i], (*arr)[high] = (*arr)[high], (*arr)[i]
	return arr, i
}

func quickSortHands(arr *[]Hand, low, high int) *[]Hand {
	fmt.Println("BEFPRE", *arr)
	sort.Slice((*arr)[:], func(i, j int) bool {

		return !((*arr)[i].strength > (*arr)[j].strength || (*arr)[i].strength == (*arr)[j].strength && CompareTwoElements((*arr)[i].cards, (*arr)[j].cards) == -1)
	})
	fmt.Println("AFTER", *arr)
	// if low < high {
	// 	var p int
	// 	arr, p = partition(arr, low, high)
	// 	arr = quickSortHands(arr, low, p-1)
	// 	arr = quickSortHands(arr, p+1, high)
	// }

	return arr
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

		hands = append(hands, Hand{bid: bid, cards: result[0], strength: CalculateStrength(result[0])})

		sizeOfHands := len(hands)

		if sizeOfHands == 1 {
			continue
		}
	}

	hands = *quickSortHands(&hands, 0, len(hands)-1)

	var result = 0

	for strength, hand := range hands {
		fmt.Println(hand.cards, hand.bid)
		result += hand.bid * (strength + 1)
	}

	fmt.Println(result)

	if result != 6839 {
		fmt.Println("Result should be 6839!")
	}
}

// 244220363 not correct
// 244226035 not correct
// 241656503 not correct
// 241813072 not correct

// short answer correct 6839 - alternative solution
