package main

import (
	"advent_helper/file_loader"
	"bufio"
	"fmt"
	"regexp"
)

type Node struct {
	value string
	L     *Node
	R     *Node
}

var regexMatcher = regexp.MustCompile(`[A-Z]{3}`)

var left byte = byte('L')
var right byte = byte('R')
var endLetter byte = byte('Z')

func createNode(value string) *Node {
	return &Node{value: value, L: nil, R: nil}
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

	return &nodeMap, &heads, instructions
}

func HeadsEndWithLetter(heads *[](*Node), value byte) bool {
	var result bool = true

	for _, head := range *heads {
		if head.value[len(head.value)-1] != value {
			result = false
		}
	}

	return result
}

func moveHeads(heads *[](*Node), instruction byte) {
	for headIndex := 0; headIndex < len(*heads); headIndex++ {
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
		moveHeads(heads, instructions[i])

		printHeads(heads)

		if HeadsEndWithLetter(heads, endLetter) {
			return i + stepCount + 1
		}

		if i == len(instructions)-1 {
			stepCount += i + 1
			i = -1

			fmt.Println("stepCount", stepCount)
		}
	}

	return stepCount + 1
}

func main() {
	nodeMap, heads, instructions := ExtractNodeMapFromFile("./input.txt")

	fmt.Println(FindNode(heads, instructions, 0, nodeMap))
}
