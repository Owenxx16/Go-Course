package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// Phương thức có receiver là biến `a` thuộc kiểu `authenticationInfo`
func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func main() {
	auth := authenticationInfo{
		username: "textio_user",
		password: "securepass123",
	}

	fmt.Println(auth.getBasicAuth())
	// Output: Authorization: Basic textio_user:securepass123
}
