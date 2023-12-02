package main_test

import (
	"fmt"
	"regexp"
	"testing"
)

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func TestSomething(t *testing.T) {
	r, _ := regexp.Compile("(?=(\\d|one|two|three))")
	str := "twone"
	matches := r.FindAllString(str, -1)
	fmt.Println(matches) // Outputs: [a a a]
}
