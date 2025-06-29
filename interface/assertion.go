package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {

	em, em1 := e.(email)
	sm, sm1 := e.(sms)
	if em1 {
		return em.toAddress, em.cost()
	}
	if sm1 {
		return sm.toPhoneNumber, sm.cost()
	}
	return "", 0.0
	// ?
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	e1 := email{isSubscribed: false, body: "Hello", toAddress: "a@example.com"}
	e2 := sms{isSubscribed: true, body: "Hi there", toPhoneNumber: "123456789"}
	e3 := invalid{}

	fmt.Println(getExpenseReport(e1)) // Output: a@example.com 0.25
	fmt.Println(getExpenseReport(e2)) // Output: 123456789 0.21
	fmt.Println(getExpenseReport(e3)) // Output: "" 0
}
