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

var nodeMap = make(map[string]*Node)

var regexMatcher = regexp.MustCompile(`[A-Z]{3}`)

func main() {
	readFile := file_loader.OpenFile("./input_short_short.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	var head *Node = nil

	// Get the instructions
	fileScanner.Scan()
	fileScanner.Text()

	// Step over an empty line
	fileScanner.Scan()
	// Gets all the nodes
	for {
		fileScanner.Scan()
		line := fileScanner.Text()

		if line == "" {
			break
		}

		var matches = regexMatcher.FindAllString(line, -1)
		fmt.Println(regexMatcher.FindAllString(line, -1))

		_, ok := nodeMap[matches[0]]

		if !ok {
			nodeMap[matches[0]] = &Node{value: matches[0], L: nil, R: nil}

			fmt.Println(nodeMap[matches[0]])
		}

		_, ok1 := nodeMap[matches[1]]

		if !ok1 {
			nodeMap[matches[1]] = &Node{value: matches[1], L: nil, R: nil}
			fmt.Println(nodeMap[matches[1]])
		}

		_, ok2 := nodeMap[matches[2]]

		if !ok2 {
			nodeMap[matches[2]] = &Node{value: matches[2], L: nil, R: nil}
			fmt.Println(nodeMap[matches[2]])
		}

		if head == nil {
			head = nodeMap[matches[0]]
		}

		nodeMap[matches[0]].L = nodeMap[matches[1]]
		nodeMap[matches[0]].R = nodeMap[matches[2]]
	}

	for _, item := range nodeMap {
		fmt.Println(item.value, item.L.value, item.R.value)
	}
}
