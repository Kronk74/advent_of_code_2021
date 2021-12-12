package days

import "strconv"

func Day1Part1(input []int) string {
	counter := 0

	for c := 1; c < len(input); c++ {
		if input[c-1] < input[c] {
			counter++
		}
	}

	return strconv.Itoa(counter)
}

func Day1Part2(input string) string {
	result := "nope"

	return result
}
