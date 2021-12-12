package days

import (
	"fmt"
	"os"

	days "github.com/Kronk74/advent_of_code_2021/days/day1"
	"github.com/Kronk74/advent_of_code_2021/utils/aocg"
	"github.com/Kronk74/advent_of_code_2021/utils/files"
)

func CallDay(day int, part int) string {
	path, err := os.Getwd()
	aocg.Check(err)
	var result string

	switch {
	//day1
	case day == 1 && part == 1:
		dayFolderPath := fmt.Sprint(path, "/days/day", day, "/input")
		result = days.Day1Part1(files.GetInputInteger(dayFolderPath))

	case day == 1 && part == 2:
		dayFolderPath := fmt.Sprint(path, "/days/day", day, "/input")
		result = days.Day1Part2(dayFolderPath)

	}

	return result
}
