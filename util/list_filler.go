/*Package util holds utility functions.
 * This file holds functions to fill inner lists from outer files (see lists directory).
 * This functions used to load task lists to map or slice of values.
 */
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

	// the file breaks by the lines
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		list[line[0]] = strings.Split(line[1], ",")
	}

	return list
}

// FillList collect usual list information from provided file
func FillList(sourceFile string) []string {
	var list []string

	file, err := os.Open(sourceFile)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		list = append(list, strings.Split(scanner.Text(), ",")...)
	}

	return list
}
