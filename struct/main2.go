package main

import (
	"fmt"
	"testing"
)

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	}
	return "", false
}

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

// i did it
func newUser(name string, membershipType string) User {
	// ?
	if membershipType == "premium" {
		return User{
			Name: name,
			Membership: Membership{
				Type:             membershipType,
				MessageCharLimit: 1000,
			},
		}
	} else {
		return User{
			Name: name,
			Membership: Membership{
				Type:             membershipType,
				MessageCharLimit: 100,
			},
		}
	}
}

// fix

// func newUser(name string, membershipType string) User {
// 	membership := Membership{Type: membershipType}
// 	if membershipType == "premium" {
// 		membership.MessageCharLimit = 1000
// 	} else {
// 		membership.Type = "standard"
// 		membership.MessageCharLimit = 100
// 	}
// 	return User{Name: name, Membership: membership}
// }

func Test(t *testing.T) {
	type testCase struct {
		name           string
		membershipType string
	}

	runCases := []testCase{
		{"Syl", "standard"},
		{"Pattern", "premium"},
		{"Pattern", "standard"},
	}

	submitCases := append(runCases, []testCase{
		{"Renarin", "standard"},
		{"Lift", "premium"},
		{"Dalinar", "standard"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		user := newUser(test.name, test.membershipType)

		msgCharLimit := 100
		if test.membershipType == "premium" {
			msgCharLimit = 1000
		}

		if user.Name != test.name {
			failCount++
			t.Errorf(`---------------------------------
Test Failed (name):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, test.name, test.membershipType, test.name, user.Name)
		} else if user.Membership.Type != test.membershipType {
			failCount++
			t.Errorf(`---------------------------------
Test Failed (membership type):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, test.name, test.membershipType, test.membershipType, user.Membership.Type)
		} else if user.Membership.MessageCharLimit != msgCharLimit {
			failCount++
			t.Errorf(`---------------------------------
Test Failed (message character limit):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, test.name, test.membershipType, msgCharLimit, user.Membership.MessageCharLimit)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v, %v, %v
Actual:     %v, %v, %v
`, test.name, test.membershipType, test.name, test.membershipType, msgCharLimit, user.Name, user.Membership.Type, user.Membership.MessageCharLimit)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

func main() {
	// Example: run the same cases and print results
	cases := []struct {
		name, membershipType string
	}{
		{"Syl", "standard"},
		{"Pattern", "premium"},
		{"Pattern", "standard"},
	}
	for _, tc := range cases {
		u := newUser(tc.name, tc.membershipType)
		fmt.Printf("%+v\n", u)
	}
}
