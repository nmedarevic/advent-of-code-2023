package matrix

func ConvertLinesToMatrix(lines []string) [][]string {
	var matrix = make([][]string, len(lines))

	for i, line := range lines {
		matrix[i] = make([]string, len(line))

		for j, char := range line {
			matrix[i][j] = string(char)
		}
	}

	return matrix
}
