package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	numberStr := strconv.Itoa(n)
	sum := float64(0)

	for _, v := range numberStr {
		digit, _ := strconv.Atoi(string(v))
		sum = sum + math.Pow(float64(digit), float64(len(numberStr)))
	}

	if sum == float64(n) {
		return true
	}

	return false
}