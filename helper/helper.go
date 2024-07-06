package helper

import "strings"

func CountNonEmptyLines(input []byte) int {
	var count int
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		switch line {
		case "", "\n":
		default:
			count++
		}

	}
	return count
}
