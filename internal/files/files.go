package files

import (
	"os"
	"strings"
)

func readFile(fileName string) (string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	newString := string(file[:])
	return newString, nil
}

func getLastIndex(fileName string) (int, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return 0, nil
	}
	newString := string(file)
	rows := strings.Split(newString, "\n")
	return len(rows), nil
}
