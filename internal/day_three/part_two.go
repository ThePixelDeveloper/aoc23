package day_three

import (
	"strconv"
	"strings"
)

//go:noinline
func partTwo(input []byte) int {
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

	var stars map[int]map[int][]int

	stars = make(map[int]map[int][]int)

	for _, number := range numbers {
		for i := number.row - 1; i <= number.row+1; i++ {
			for j := number.left - 1; j <= number.right+1; j++ {
				if rows[i][j] == "*" {
					num, _ := strconv.Atoi(number.number)
					if _, ok := stars[i]; !ok {
						stars[i] = make(map[int][]int)
					}
					stars[i][j] = append(stars[i][j], num)
					continue
				}
			}
		}
	}

	var output int

	for _, x := range stars {
		for _, y := range x {
			if len(y) != 2 {
				continue
			}

			output += y[0] * y[1]
		}
	}

	return output
}
