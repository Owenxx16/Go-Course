package main

func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
	nameCounts := make(map[rune]map[string]int)
	for _, name := range names {
		firstChar := rune(name[0])
		if _, exists := nameCounts[firstChar]; exists {
			nameCounts[firstChar] = make(map[string]int)
		}
		nameCounts[firstChar][name]++
	}
	return nameCounts
}
