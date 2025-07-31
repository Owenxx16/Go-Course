package main

func bulkSend(numMessages int) float64 {
	// ?
	var total = 0.0;
	for i := 1; i < numMessages; i++ {
		sum := float64(i-1) * 0.01;
		total += 1.0 + sum
}
	return total
}
