package main

func addEmailsToQueue(emails []string) chan string {
	// ?
	emailQueue := make(chan string, len(emails))
	for _, email := range emails {
		emailQueue <- email
	}
	close(emailQueue)
	return emailQueue
}
