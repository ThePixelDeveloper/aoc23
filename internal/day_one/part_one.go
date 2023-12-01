package day_one

//go:noinline
func partOne(input []byte) int {
	var total int
	var value int

	// Ten's
	for i := 0; i < len(input); i++ {
		// From 1 to 9 in ASCII
		if value == 0 && input[i] >= 49 && input[i] <= 57 {
			value = int(input[i])<<3 + int(input[i])<<1
			total += value - 480
			continue
		}

		// Check for newlines
		if input[i] == 10 {
			value = 0
		}
	}

	value = 0

	// Unit's
	for i := len(input) - 1; i > 0; i-- {
		// From 1 to 9 in ASCII
		if value == 0 && input[i] >= 49 && input[i] <= 57 {
			value = int(input[i])
			total += value - 48
			continue
		}

		// Check for newlines
		if input[i] == 10 {
			value = 0
		}
	}

	return total
}
