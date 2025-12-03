package wordcount

import (
	"fmt"
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	wordFrequency := Frequency{}

	re := regexp.MustCompile("[,.;:\n\t!@#$%^& ]+")
	parts := (re.Split(phrase, -1))


	for i, v := range parts {
		if strings.HasPrefix(v, "'") && strings.HasSuffix(v, "'") {
			parts[i] = strings.Trim(v, "'")
		}
		if strings.HasPrefix(v, "'") || strings.HasSuffix(v, "'") {
			parts[i] = strings.Trim(v, "'")
		}


		if v == " "{
			continue
		}
	}

	fmt.Println(parts)

	new := strings.ToLower(strings.Join(parts, " "))
	parts = strings.Split(new, " ")
	
	for _, v := range parts {
		_, ok := wordFrequency[v]

		if v == ""{
			continue
		}

		if ok {
			wordFrequency[v] += 1
		} else {
			wordFrequency[v] = 1
		}
	}

	return wordFrequency
}