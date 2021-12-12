package files

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/Kronk74/advent_of_code_2021/utils/aocg"
)

func GetInputString(input string) []string {
	var result []string

	file, err := os.Open(input)
	if err != nil {
		aocg.Check(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetInputInteger(input string) []int {
	var result []int

	file, err := os.Open(input)
	if err != nil {
		aocg.Check(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r, _ := strconv.Atoi(scanner.Text())
		result = append(result, r)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
