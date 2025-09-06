package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	if err := em.recipient.updateStatus("email_bounced"); err != nil {
		return fmt.Errorf("error updating user status: %w", err)
	}
	if err := a.track("email_bounced"); err != nil {
		return fmt.Errorf("error tracking user bounce: %w", err)
	}
	return nil
}
