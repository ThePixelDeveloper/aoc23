package day_two

//go:noinline
func partTwo(input []byte) int {
	var maxRed, maxGreen, maxBlue, red, green, blue, total, game int

	for i := 0; i < len(input); i++ {
		// G
		if input[i] == 71 {

			// Reset game state
			maxRed = 0
			maxGreen = 0
			maxBlue = 0
			red = 0
			green = 0
			blue = 0

			// We know it's Game, so we can skip ahead to find ':'
			for input[i] != 58 {
				i++
			}

			i--

			// Now reverse to build the game number
			if input[i] >= 48 && input[i] <= 57 {
				game += int(input[i]) - 48
			}

			if input[i-1] >= 48 && input[i-1] <= 57 {
				game += (int(input[i-1])<<3 + int(input[i-1])<<1) - 480
			}

			if input[i-2] >= 48 && input[i-2] <= 57 {
				game += (int(input[i-2])<<6 + int(input[i-2])<<5 + int(input[i-2])<<2) - 4800
			}
		}

		// r
		if input[i] == 32 && input[i+1] == 114 {
			i--

			// Now reverse to build the red number
			if input[i] >= 48 && input[i] <= 57 {
				red += int(input[i]) - 48
			}

			if input[i-1] >= 48 && input[i-1] <= 57 {
				red += (int(input[i-1])<<3 + int(input[i-1])<<1) - 480
			}

			i += 4
		}

		// g
		if input[i] == 32 && input[i+1] == 103 {
			i--

			// Now reverse to build the green number
			if input[i] >= 48 && input[i] <= 57 {
				green += int(input[i]) - 48
			}

			if input[i-1] >= 48 && input[i-1] <= 57 {
				green += (int(input[i-1])<<3 + int(input[i-1])<<1) - 480
			}

			i += 6
		}

		// b
		if input[i] == 32 && input[i+1] == 98 {
			i--

			// Now reverse to build the blue number
			if input[i] >= 48 && input[i] <= 57 {
				blue += int(input[i]) - 48
			}

			if input[i-1] >= 48 && input[i-1] <= 57 {
				blue += (int(input[i-1])<<3 + int(input[i-1])<<1) - 480
			}

			i += 5
		}

		// Semicolon / Newline
		if input[i] == 59 || input[i] == 10 {
			if red > maxRed {
				maxRed = red
			}
			if green > maxGreen {
				maxGreen = green
			}
			if blue > maxBlue {
				maxBlue = blue
			}
			red = 0
			green = 0
			blue = 0
		}

		if input[i] == 10 {
			total += maxRed * maxGreen * maxBlue

			game = 0
		}
	}

	if red > maxRed {
		maxRed = red
	}
	if green > maxGreen {
		maxGreen = green
	}
	if blue > maxBlue {
		maxBlue = blue
	}

	total += maxRed * maxGreen * maxBlue

	return total
}
