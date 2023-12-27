package main

import (
	"testing"
)

func TestShortOne(t *testing.T) {
	nodeMap, heads, instructions := ExtractNodeMapFromFile("./input_short_short.txt")

	numberOfSteps := FindNode(heads, instructions, 0, nodeMap)

	if numberOfSteps != 2 {
		t.Error("Wrong number of steps for short input", numberOfSteps)
	}
}

func TestShortLongerOne(t *testing.T) {
	nodeMap, head, instructions := ExtractNodeMapFromFile("./input_short_longer.txt")

	numberOfSteps := FindNode(head, instructions, 0, nodeMap)

	if numberOfSteps != 6 {
		t.Error("Wrong number of steps for longer short input", numberOfSteps)
	}
}

func TestShortWithNumbers(t *testing.T) {
	nodeMap, head, instructions := ExtractNodeMapFromFile("./input_short_with_numbers.txt")

	numberOfSteps := FindNode(head, instructions, 0, nodeMap)

	if numberOfSteps != 6 {
		t.Error("Wrong number of steps for longer short input", numberOfSteps)
	}
}

func TestHeadsEndWithLetter(t *testing.T) {
	var heads [](*Node) = [](*Node){&Node{value: "3AA"}, &Node{value: "DAA"}, &Node{value: "EAA"}}

	if !HeadsEndWithLetter(&heads, byte('A')) {
		t.Error("Not all heads point to A")
	}
}
