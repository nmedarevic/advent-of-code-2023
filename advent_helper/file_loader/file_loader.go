package file_loader

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(filePath string) *os.File {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)

		panic(err)
	}

	return readFile
}

func readLines(readFile *os.File) []string {
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}

func LoadLinesFromFile(filePath string) []string {
	readFile := openFile(filePath)

	var fileLines = readLines(readFile)

	return fileLines
}
