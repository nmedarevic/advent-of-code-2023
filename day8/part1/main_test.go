package main

import (
	"testing"
)

func TestShortOne(t *testing.T) {
	nodeMap, head, instructions := ExtractNodeMapFromFile("./input_short_short.txt")

	numberOfSteps := FindNode(head, instructions, 0, nodeMap)

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
