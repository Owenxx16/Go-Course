package main

import "errors"

type userWithDeletion struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]userWithDeletion, name string) (deleted bool, err error) {
	// ?
	elem, ok := users[name]
	if !ok {
		return false, errors.New("Not Found")
	}
	if !elem.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}
