package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	// ?
	var messages = [3]string{primary, secondary, tertiary};
	var originals = len(primary);
	var originals2 = originals + len(secondary);
	var originals3 = originals2 + len(tertiary);
	var intervals = [3] int{originals, originals2, originals3};
	return messages, intervals
}
