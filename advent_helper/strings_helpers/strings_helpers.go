package strings_helpers

import "strconv"

func ElementExistsInStringArray(array []string, element string) bool {
	for _, arrayElement := range array {
		if arrayElement == element {
			return true
		}
	}

	return false
}

func StringToNumberDefaultToZero(numberString string) int {
	number, err := strconv.Atoi(numberString)

	if err != nil {
		number = 0
	}

	return number
}

func StringToNumberAndPanic(numberString string) int {
	number, err := strconv.Atoi(numberString)

	if err != nil {
		panic("Could not convert string to number")
	}

	return number
}

func StringToNumberAndPanic64(numberString string) int64 {
	number, err := strconv.ParseInt(numberString, 10, 64)

	if err != nil {
		panic("Could not convert string to number")
	}

	return number
}
