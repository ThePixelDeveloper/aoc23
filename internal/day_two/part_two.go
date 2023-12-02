package day_two

import (
	"regexp"
	"strconv"
	"strings"
)

//go:noinline
func partTwo(input []byte) int {
	var total int
	var limit map[string]int

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {

		limit = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		//game, _ := strconv.Atoi(line[5:strings.Index(line, ":")])
		sets := strings.Split(line[strings.Index(line, ":")+2:], ";")

		for _, set := range sets {
			drawsRe := regexp.MustCompile("(([0-9]+) (blue|green|red))")
			draws := drawsRe.FindAllStringSubmatch(set, -1)
			for _, draw := range draws {
				num, _ := strconv.Atoi(draw[2])
				col := draw[3]

				if limit[col] < num {
					limit[col] = num
				}
			}
		}

		total += limit["red"] * limit["green"] * limit["blue"]
	}

	return total
}
