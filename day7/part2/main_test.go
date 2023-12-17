package main

import (
	"testing"
)

func TestAAAAA(t *testing.T) {
	if CalculateStrength("AAAAA") != StrengthMap["FiveOfAKind"] {
		t.Error("AAAAA should be", StrengthMap["FiveOfAKind"])
	}
}

func TestAA8AA(t *testing.T) {
	if CalculateStrength("AA8AA") != StrengthMap["FourOfAKind"] {
		t.Error("AA8AA should be", StrengthMap["FourOfAKind"])
	}
}

func Test23332(t *testing.T) {
	if CalculateStrength("23332") != StrengthMap["FullHouse"] {
		t.Error("23332 should be", StrengthMap["FullHouse"])
	}
}

func TestTTT98(t *testing.T) {
	if CalculateStrength("TTT98") != StrengthMap["ThreeOfAKind"] {
		t.Error("TTT98 should be", StrengthMap["ThreeOfAKind"])
	}
}

func Test23432(t *testing.T) {
	if CalculateStrength("23432") != StrengthMap["TwoPair"] {
		t.Error("23432 should be", StrengthMap["TwoPair"])
	}
}

func TestA23A4(t *testing.T) {
	if CalculateStrength("A23A4") != StrengthMap["OnePair"] {
		t.Error("A23A4 should be", StrengthMap["OnePair"])
	}
}

func Test23456(t *testing.T) {
	if CalculateStrength("23456") != StrengthMap["HighCard"] {
		t.Error("23456 should be", StrengthMap["HighCard"])
	}
}

func Test32T3K(t *testing.T) {
	if CalculateStrength("32T3K") != StrengthMap["OnePair"] {
		t.Error("32T3K should be ", StrengthMap["OnePair"])
	}
}

func TestKK677(t *testing.T) {
	if CalculateStrength("KK677") != StrengthMap["TwoPair"] {
		t.Error("KK677 should be ", StrengthMap["TwoPair"])
	}
}

func TestT55J5(t *testing.T) {
	if CalculateStrength("T55J5") != StrengthMap["FourOfAKind"] {
		t.Error("T55J5 should be ", StrengthMap["FourOfAKind"])
	}
}

func TestKTJJT(t *testing.T) {
	if CalculateStrength("KTJJT") != StrengthMap["FourOfAKind"] {
		t.Error("KTJJT should be ", StrengthMap["FourOfAKind"])
	}
}

func TestQQQJA(t *testing.T) {
	if CalculateStrength("QQQJA") != StrengthMap["FourOfAKind"] {
		t.Error("QQQJA should be ", StrengthMap["FourOfAKind"])
	}
}
