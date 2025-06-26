package main

import (
	"fmt"
	"reflect"
)

type contact struct {
	sendingLimit int32
	userID       string
	age          int32
}

type perms struct {
	canSend         bool
	canReceive      bool
	permissionLevel int
	canManage       bool
}

func main() {
	checkContact := reflect.TypeOf(contact{})
	checkPerms := reflect.TypeOf(perms{})
	fmt.Println("Contact struct: %d ", checkContact.Size())
	fmt.Println("Permissions struct: %d ", checkPerms.Size())
}
