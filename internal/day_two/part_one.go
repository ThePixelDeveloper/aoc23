package day_two

import (
	"regexp"
	"strconv"
	"strings"
)

//go:noinline
func partOne(input []byte) int {
	var total int
	var limit map[string]int
	var valid bool

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		limit = map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		}

		valid = true

		game, _ := strconv.Atoi(line[5:strings.Index(line, ":")])
		sets := strings.Split(line[strings.Index(line, ":")+2:], ";")

		for _, set := range sets {
			drawsRe := regexp.MustCompile("(([0-9]+) (blue|green|red))")
			draws := drawsRe.FindAllStringSubmatch(set, -1)
			for _, draw := range draws {
				num, _ := strconv.Atoi(draw[2])
				col := draw[3]

				if limit[col]-num < 0 {
					valid = false
				}
			}
		}

		if valid {
			total += game
		}
	}

	return total
}
