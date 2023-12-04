package util

import (
	"bufio"
	"fmt"
	"os"
)

var file *os.File

func GetScanner(day int) *bufio.Scanner {
	filePath := fmt.Sprintf("./inputs/input%d.txt", day)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return scanner
}

func CloseScanner() {
	file.Close()
}
