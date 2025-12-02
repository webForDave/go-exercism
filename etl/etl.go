package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {

	transformed := map[string]int{} 

	for k, v := range in {
		for _, char := range v {
			transformed[strings.ToLower(char)] = k
		}
	}

	return transformed
}