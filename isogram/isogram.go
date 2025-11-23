package isogram

import (
	"strings"
)

func IsIsogram(word string) bool{

	for _, character := range word{
		if string(character) == "-" || string(character) == " "{
			continue
		}
		if strings.Count(strings.ToLower(word), strings.ToLower(string(character))) > 1{
			return false
		}
	}

	return true
}