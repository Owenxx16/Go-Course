package main

import "fmt"

func (e email) cost() int {
	var total int
	if e.isSubscribed {
		total = 2 * len(e.body)
	} else {
		total = 5 * len(e.body)
	}
	return total
	// ?
}

func (e email) format() string {
	// ?
	var prefix string
	if e.isSubscribed {
		prefix = fmt.Sprintf("'%s' | Subscribed", e.body)
	} else {
		prefix = fmt.Sprintf("'%s' | Not Subscribed", e.body)
	}
	return prefix
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

func main() {
	e := email{isSubscribed: false, body: "Hello"}
	fmt.Println(e.cost())   // Output: 25
	fmt.Println(e.format()) // Output: 'Hello' | Not Subscribed
}
