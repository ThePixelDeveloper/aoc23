package day_three

import (
	"strconv"
	"strings"
)

//go:noinline
func partOne(input []byte) int {
	// Add padding for easier check
	lineLen := len(strings.Split(string(input), "\n")[0])

	// Top and Bottom padding
	puzzle := strings.Repeat(".", lineLen) + "\n" + string(input) + "\n" + strings.Repeat(".", lineLen) + "\n"

	var rows [][]string

	for _, row := range strings.Split(puzzle, "\n") {
		var cols []string

		for _, col := range strings.Split("."+row+".", "") {
			cols = append(cols, col)
		}

		rows = append(rows, cols)
	}

	var numbers []struct {
		number string
		row    int
		left   int
		right  int
	}

	for i, row := range rows {
		var number string
		var left int

		for j, col := range row {
			if _, err := strconv.Atoi(col); err == nil {
				if number == "" {
					left = j
				}
				number = number + col
			} else if number != "" {
				numbers = append(numbers, struct {
					number string
					row    int
					left   int
					right  int
				}{number: number, row: i, left: left, right: j - 1})
				number = ""
				left = 0
			}
		}
	}

	var output int

	for _, number := range numbers {
		for _, s := range []string{"&", "%", "@", "/", "=", "-", "*", "+", "#", "$"} {
			for i := number.row - 1; i <= number.row+1; i++ {
				for j := number.left - 1; j <= number.right+1; j++ {
					if rows[i][j] == s {
						num, _ := strconv.Atoi(number.number)
						output += num
						continue
					}
				}
			}
		}
	}

	return output
}
