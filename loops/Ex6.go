package main

func countConnections(groupSize int) int {
	// ?
	var connections = 0;
	for i := 1; i < groupSize; i++ {
		connections += i
	}
	return connections
}
