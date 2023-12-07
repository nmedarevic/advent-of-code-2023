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
