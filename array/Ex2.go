package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	switch plan {
	case "pro plan":
		return messages[:], nil
	case "free plan":
		return messages[:2], nil
	default:
		return nil, errors.New("unsupported plan")
	}
}
