package main

import (
	"errors"
)

func validateStatus(status string) error {
	//
	if len(status) > 0 && len(status) <= 140 {
		return errors.New("status must be a non-empty string with a maximum length of 140 characters");
	} 
	return nil
}
