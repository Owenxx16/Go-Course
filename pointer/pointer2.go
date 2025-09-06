package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message1 struct {
	Recipient string
	Success   bool
}

// don't touch above this line

// ?
func getMessageText1(ana *Analytics, m Message1) bool {
	var boola = false
	ana.MessagesTotal++
	if m.Success {
		ana.MessagesSucceeded++
		boola = true
	} else {
		ana.MessagesFailed++
		boola = false
	}
	return boola
}
