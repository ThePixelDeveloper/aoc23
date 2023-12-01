package day_one

//go:noinline
func partTwo(input []byte) int {
	var total int
	var value int

	// Written numbers (Ten's)
	for i := 0; i < len(input); i++ {
		if input[i] == 10 {
			value = 0
		} else if value == 0 {
			if input[i] == 111 && input[i+2] == 101 && input[i+1] == 110 { // One
				value = 1
				total += (value << 3) + (value << 1)
			} else if input[i] == 116 && input[i+2] == 111 && input[i+1] == 119 { // Two
				value = 2
				total += (value << 3) + (value << 1)
			} else if input[i] == 115 && input[i+2] == 120 && input[i+1] == 105 { // Six
				value = 6
				total += (value << 3) + (value << 1)
			} else if input[i] == 102 && input[i+3] == 114 && input[i+1] == 111 && input[i+2] == 117 { // Four
				value = 4
				total += (value << 3) + (value << 1)
			} else if input[i] == 110 && input[i+3] == 101 && input[i+1] == 105 && input[i+2] == 110 { // Nine
				value = 9
				total += (value << 3) + (value << 1)
			} else if input[i] == 102 && input[i+3] == 101 && input[i+1] == 105 && input[i+2] == 118 { // Five
				value = 5
				total += (value << 3) + (value << 1)
			} else if input[i] == 116 && input[i+4] == 101 && input[i+1] == 104 && input[i+2] == 114 && input[i+3] == 101 { // Three
				value = 3
				total += (value << 3) + (value << 1)
			} else if input[i] == 115 && input[i+4] == 110 && input[i+1] == 101 && input[i+2] == 118 && input[i+3] == 101 { // Seven
				value = 7
				total += (value << 3) + (value << 1)
			} else if input[i] == 101 && input[i+4] == 116 && input[i+1] == 105 && input[i+2] == 103 && input[i+3] == 104 { // Eight
				value = 8
				total += (value << 3) + (value << 1)
			} else if input[i] >= 49 && input[i] <= 57 {
				value = int(input[i])<<3 + int(input[i])<<1
				total += value - 480
			}
		}
	}

	value = 0

	// Written numbers (Units's)
	for i := len(input) - 1; i > 0; i-- {
		if input[i] == 10 {
			value = 0
		} else if value == 0 {
			if input[i] == 101 && input[i-2] == 111 && input[i-1] == 110 { // One
				value = 1
				total += value
			} else if input[i] == 111 && input[i-2] == 116 && input[i-1] == 119 { // Two
				value = 2
				total += value
			} else if input[i] == 120 && input[i-2] == 115 && input[i-1] == 105 { // Six
				value = 6
				total += value
			} else if input[i] == 114 && input[i-3] == 102 && input[i-1] == 117 && input[i-2] == 111 { // Four
				value = 4
				total += value
			} else if input[i] == 101 && input[i-3] == 110 && input[i-1] == 110 && input[i-2] == 105 { // Nine
				value = 9
				total += value
			} else if input[i] == 101 && input[i-3] == 102 && input[i-1] == 118 && input[i-2] == 105 { // Five
				value = 5
				total += value
			} else if input[i] == 101 && input[i-4] == 116 && input[i-1] == 101 && input[i-2] == 114 && input[i-3] == 104 { // Three
				value = 3
				total += value
			} else if input[i] == 110 && input[i-4] == 115 && input[i-1] == 101 && input[i-2] == 118 && input[i-3] == 101 { // Seven
				value = 7
				total += value
			} else if input[i] == 116 && input[i-4] == 101 && input[i-1] == 104 && input[i-2] == 103 && input[i-3] == 105 { // Eight
				value = 8
				total += value
			} else if input[i] >= 49 && input[i] <= 57 {
				value = int(input[i])
				total += value - 48
			}
		}
	}

	return total
}
