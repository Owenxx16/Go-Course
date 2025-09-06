package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	// ?
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
		default:
			waitForData(logChan)
			time.Sleep(1 * time.Second) // Avoid busy waiting
		}
	}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
