package array_helpers

func FindIntersection(array1 *[]int, array2 *[]int) *[]int {
	var longerArray = array1
	var shorterArray = array2

	if len(*array2) > len(*array1) {
		longerArray = array2
		shorterArray = array1
	}

	var result []int

	for _, item := range *longerArray {
		for _, item2 := range *shorterArray {
			if item == item2 {
				result = append(result, item)
			}
		}
	}

	return &result
}

func FindLowestNumber(numberArray []int) int {
	lowestNumber := numberArray[0]

	for _, number := range numberArray {
		if number < lowestNumber {
			lowestNumber = number
		}
	}

	return lowestNumber
}

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

func FindMin[T Number](numberArray []T) T {
	lowestNumber := numberArray[0]

	for _, number := range numberArray {
		if number < lowestNumber {
			lowestNumber = number
		}
	}

	return lowestNumber
}
