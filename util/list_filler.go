package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// FillTaskList creates and fills list with provided file data
func FillTaskList(sourceFile string) map[string][]string {
	list := make(map[string][]string, 0)

	file, err := os.Open(sourceFile)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		list[line[0]] = strings.Split(line[1], ",")
	}

	return list
}
