package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	output := []string{}

	for i := range rhyme {
		if i == len(rhyme) - 1 {
			output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		} else {
			output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i + 1]))
		}
	}

	return output
}