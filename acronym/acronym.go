package acronym

import (
	"strings"
)

func Split(r rune) bool {
	return r == '-' || r == ' ' || r == '_'
}

func Abbreviate(s string) string {
	words := strings.FieldsFunc(strings.ToUpper(s), Split)
	var result strings.Builder

	for _, v := range words {
		result.WriteString(string(v[0]))
	}

	return result.String()
}