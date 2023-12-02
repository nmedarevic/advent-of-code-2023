package strings

func ElementExistsInStringArray(array []string, element string) bool {
	for _, arrayElement := range array {
		if arrayElement == element {
			return true
		}
	}

	return false
}
