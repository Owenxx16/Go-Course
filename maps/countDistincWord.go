package main

import "strings"

func countDistinctWords(messages []string) int {
	// ?
	wordSet := make(map[string]struct{})
	for _, message := range messages {
		words := strings.Split(message, " ")
		for _, word := range words {
			normalized = strings.ToLower(word)
			wordSet[normalized] = struct{}{}
		}
	}
	return len(wordSet)
}
