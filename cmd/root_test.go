package cmd

import "testing"

func TestMainFunc(t *testing.T) {
	var (
		message  string
		expected string
		name     string
	)
	name = "Go"
	message = getMessage(name)
	expected = "Hello Go!"
	if message != expected {
		t.Errorf("Main returned %s, expected %s", message, expected)
	}

	name = "Gopher"
	message = getMessage(name)
	expected = "Hello Gopher!"
	if message != expected {
		t.Errorf("Main returned %s, expected %s", message, expected)
	}
}
