package main

import "fmt"

func temp() {
	var message string
	message = getMessage()

	fmt.Println(message)
}

func getMessage() string {
	return "Hello Go!"
}
