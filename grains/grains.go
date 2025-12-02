package grains

import (
	"errors"
)

// Calculates the number of grains in the square n
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0,  errors.New("invalid input")
	}

	return 1 << (n - 1), nil
}

// Calculates the total number of grains on a chess board
func Total() uint64 {
	return 1 << 64 - 1
}