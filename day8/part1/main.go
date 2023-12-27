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
var endString string = "ZZZ"

func createNode(value string) *Node {
	return &Node{value: value, L: nil, R: nil}
}

func upsertNodeToMap(value string, nodeMap *map[string]*Node) {
	_, ok := (*nodeMap)[value]

	if !ok {
		(*nodeMap)[value] = createNode(value)
	}
}

func ExtractNodeMapFromFile(filePath string) (*map[string]*Node, *Node, string) {
	readFile := file_loader.OpenFile(filePath)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	var head *Node = nil

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

		if head == nil {
			head = nodeMap[matches[0]]
		}

		nodeMap[matches[0]].L = nodeMap[matches[1]]
		nodeMap[matches[0]].R = nodeMap[matches[2]]
	}

	return &nodeMap, head, instructions
}

func FindNode(head *Node, instructions string, stepCount int, nodeMap *map[string]*Node) int {
	if head.value == endString {
		return stepCount
	}

	var left = byte('L')
	var right = byte('R')

	for i := 0; i < len(instructions); i++ {
		if instructions[i] == left {
			head = head.L
		} else if instructions[i] == right {
			head = head.R
		}

		// fmt.Println(head.value)

		if head.value == endString {
			return i + stepCount + 1
		}

		if i == len(instructions)-1 {

			stepCount += i + 1
			fmt.Println("AGAIN", stepCount)
			i = -1
		}
	}

	return stepCount + 1
}

func main() {
	nodeMap, head, instructions := ExtractNodeMapFromFile("./input.txt")

	fmt.Println(FindNode(head, instructions, 0, nodeMap))
}
