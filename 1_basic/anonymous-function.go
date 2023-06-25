package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {

	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	blacklist := func(nama string) bool {
		return nama == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("rida", blacklist)
}
