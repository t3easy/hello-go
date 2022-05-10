package main

import "fmt"

func main() {
	var message string
	message = getMessage()

	fmt.Println(message)
}

func getMessage() string {
	return "Hello Go!"
}
