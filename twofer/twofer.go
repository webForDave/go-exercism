// Package twofer implements the sharing of cookies gotten on discount with another person.
package twofer 

import(
	"fmt"
)

// ShareWith allows a person to say the name of the person they share their cookie with.
// If the name of the person is unknown, it replaces the placeholder for their name with 'you'.
func ShareWith(name string) string{
	if len(name) < 1{
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}