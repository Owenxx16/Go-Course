package main

func isValidPassword(password string) bool {
	// ?
	hasUpper := false
	hasDigit := false
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= '0' && char <= '9':
			hasDigit = true
		default:
			return false
		}
	}
	return hasUpper && hasDigit
}
