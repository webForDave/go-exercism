// Package raindrops converts a number into its corresponding raindrop sounds.
package raindrops

import(
	"strconv"
)

// Convert returns the sound of raindrops based on whether 
// the number argument is evenly divisible by certain numbers (3, 5, 7).
func Convert(number int) string{
	var result string

	switch{
	case number % 3 == 0 && number % 5 == 0 && number % 7 == 0:
		result = "PlingPlangPlong"
	case number % 3 == 0 && number % 5 == 0:
		result = "PlingPlang"
	case number % 3 == 0 && number % 7 == 0:
		result = "PlingPlong"
	case number % 5 == 0 && number % 7 == 0:
		result = "PlangPlong"
	case number % 3 == 0:
		result = "Pling"
	case number % 5 == 0:
		result = "Plang"
	case number % 7 == 0:
		result = "Plong"
	default:
		result = strconv.Itoa(number)
	}

	return result
}