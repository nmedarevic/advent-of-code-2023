package main

import (
	"testing"
)

func TestAAAAA(t *testing.T) {
	if CalculateStrength("AAAAA") != 6 {
		t.Error("AAAAA should be 5")
	}
}

func TestAA8AA(t *testing.T) {
	if CalculateStrength("AA8AA") != 5 {
		t.Error("AA8AA should be 4")
	}
}

func Test23332(t *testing.T) {
	if CalculateStrength("23332") != 4 {
		t.Error("23332 should be 4")
	}
}

func TestTTT98(t *testing.T) {
	if CalculateStrength("TTT98") != 3 {
		t.Error("TTT98 should be 3")
	}
}

func Test23432(t *testing.T) {
	if CalculateStrength("23432") != 2 {
		t.Error("23432 should be 2")
	}
}

func TestA23A4(t *testing.T) {
	if CalculateStrength("A23A4") != 1 {
		t.Error("A23A4 should be 1")
	}
}

func Test23456(t *testing.T) {
	if CalculateStrength("23456") != 0 {
		t.Error("23456 should be 0")
	}
}
