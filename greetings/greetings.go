package greetings

import "fmt"

// Hello returns a greeting for the named person
// functions starting with a capital letter can be called by functions not in the same package
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	// := is the short for, var message string and then fmt.Sprintf on the next line
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
