package main

import (
	"advent_helper/file_loader"
	"bufio"
	"fmt"
	"regexp"
)

type Node struct {
	value    string
	L        *Node
	R        *Node
	endIndex int
}

var regexMatcher = regexp.MustCompile(`[A-Z]{3}`)

var left byte = byte('L')
var right byte = byte('R')
var endLetter byte = byte('Z')

func createNode(value string) *Node {
	return &Node{value: value, L: nil, R: nil, endIndex: 0}
}

func upsertNodeToMap(value string, nodeMap *map[string]*Node) {
	_, ok := (*nodeMap)[value]

	if !ok {
		(*nodeMap)[value] = createNode(value)
	}
}

func ExtractNodeMapFromFile(filePath string) (*map[string]*Node, *[]*Node, string) {
	readFile := file_loader.OpenFile(filePath)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	var heads [](*Node) = [](*Node){}

	// Get the instructions
	fileScanner.Scan()
	instructions := fileScanner.Text()

	// Step over an empty line
	fileScanner.Scan()

	var nodeMap = make(map[string]*Node)
	// Gets all the nodes
	for {
		fileScanner.Scan()
		line := fileScanner.Text()

		if line == "" {
			break
		}

		var matches = regexMatcher.FindAllString(line, -1)

		upsertNodeToMap(matches[0], &nodeMap)
		upsertNodeToMap(matches[1], &nodeMap)
		upsertNodeToMap(matches[2], &nodeMap)

		nodeMap[matches[0]].L = nodeMap[matches[1]]
		nodeMap[matches[0]].R = nodeMap[matches[2]]
	}

	// fmt.Println(nodeMap)

	for _, node := range nodeMap {
		if node.value[2] == byte('A') {
			heads = append(heads, node)
		}
	}

	// sort.SliceStable(heads, func(i, j int) bool {
	// 	return heads[i].value < heads[j].value
	// })

	return &nodeMap, &heads, instructions
}

func OneHeadEndsWithLetter(head *Node, value byte) bool {
	return (*head).value[2] == value
}

func MarkHeadEndingWithLetter(heads *[](*Node), value byte, currentStep int) {
	for _, head := range *heads {
		if head.endIndex == 0 && head.value[2] == value {
			head.endIndex = currentStep + 1
		}
	}
}

func AllHeadsFoundEnding(heads *[](*Node)) bool {
	var result bool = true

	for _, head := range *heads {
		if head.endIndex == 0 {
			result = false
		}
	}

	return result
}

func HeadsEndWithLetter(heads *[](*Node), value byte) bool {
	var result bool = true

	for _, head := range *heads {
		if head.value[2] != value {
			result = false
		}
	}

	return result
}

func moveHeads(heads *[](*Node), instruction byte) {
	for headIndex := 0; headIndex < len(*heads); headIndex++ {
		if (*heads)[headIndex].endIndex != 0 {
			continue
		}

		if instruction == left {
			(*heads)[headIndex] = (*heads)[headIndex].L
		} else if instruction == right {
			(*heads)[headIndex] = (*heads)[headIndex].R
		}
	}
}

func printHeads(heads *[](*Node)) {
	for _, head := range *heads {
		fmt.Print(head.value, " ")
	}

	fmt.Println("")
	fmt.Println("============")
}

func FindNode(heads *[](*Node), instructions string, stepCount int, nodeMap *map[string]*Node) int {
	if HeadsEndWithLetter(heads, endLetter) {
		return stepCount
	}

	for i := 0; i < len(instructions); i++ {
		// printHeads(heads)
		moveHeads(heads, instructions[i])
		// printHeads(heads)

		MarkHeadEndingWithLetter(heads, endLetter, stepCount+i)
		// if HeadsEndWithLetter(heads, endLetter) {
		// 	return i + stepCount + 1
		// }

		if AllHeadsFoundEnding(heads) {
			var integers []int = []int{}

			for _, head := range *heads {
				fmt.Println(head.value, head.endIndex)
				integers = append(integers, head.endIndex)
			}
			return LCM(integers[0], integers[1], integers[2:]...)

			break
			// fmt.Append(integers)
		}

		if i == len(instructions)-1 {
			stepCount += i + 1
			i = -1

			// printHeads(heads)
			// fmt.Println("stepCount", stepCount)
		}
	}

	return stepCount + 1
}

func main() {
	nodeMap, heads, instructions := ExtractNodeMapFromFile("./input.txt")
	printHeads(heads)
	fmt.Println(FindNode(heads, instructions, 0, nodeMap))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

//9858474970153 correct
