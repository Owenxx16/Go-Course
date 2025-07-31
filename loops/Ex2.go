package main

func maxMessages(thresh int) int {
	// ?
	var total = 0
	for i := 1; ; i++ {
		cost := 100 + (i - 1)
		if total + cost > thresh {
			return i - 1
		}
		total += cost
	}
	return total
}
