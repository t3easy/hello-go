package main

import "testing"

func TestMainFunc(t *testing.T) {
	var (
		message  string
		expected string
	)
	message = getMessage()
	expected = "Hello Go!"
	if message != expected {
		t.Errorf("Main returned %s, expected %s", message, expected)
	}
}
