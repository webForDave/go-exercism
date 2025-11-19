package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error){
	if len(a) != len(b){
		return 0, errors.New("strands length do not match")
	}

	distance := 0

	for i := range len(a){
		if string(a[i]) != string(b[i]){
			distance += 1
		}

	}
	
	return distance, nil
}