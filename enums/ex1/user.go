package main

import "fmt"

type email struct {
	status    string
	recipient *user
}

type user struct {
	email  string
	status string
}

type analytics struct {
	totalBounces int
}

func (u *user) updateStatus(status string) error {
	if status != "email_bounced" {
		return fmt.Errorf("error updating user status: %w", status)
	}
	u.status = status
	return nil
}

func (a *analytics) track(event string) error {
	if event != "email_bounced" {
		return fmt.Errorf("error tracking user bounce: %w", event)
	}
	a.totalBounces++
	return nil
}
