package bottlesong

import(
	"fmt"
	"strings"
)

var values = map[int]string {
	10: "Ten green bottles",
	9: "Nine green bottles",
	8: "Eight green bottles",
	7: "Seven green bottles",
	6: "Six green bottles",
	5: "Five green bottles",
	4: "Four green bottles",
	3: "Three green bottles",
	2: "Two green bottles",
	1: "One green bottle",
	0: "No green bottles",
}

func Recite(startBottles, takeDown int) []string {
	output := []string {}

	for i := startBottles; i > (startBottles - takeDown); i -- {
		if i - 1 ==  startBottles - takeDown {
			output = append(output, fmt.Sprintf("%v hanging on the wall,", values[i]), fmt.Sprintf("%v hanging on the wall,", values[i]), fmt.Sprintf("And if one green bottle should accidentally fall,"), fmt.Sprintf("There'll be %v hanging on the wall.", strings.ToLower(values[i - 1])))
		} else {
			output = append(output, fmt.Sprintf("%v hanging on the wall,", values[i]), fmt.Sprintf("%v hanging on the wall,", values[i]), fmt.Sprintf("And if one green bottle should accidentally fall,"), fmt.Sprintf("There'll be %v hanging on the wall.", strings.ToLower(values[i - 1])), fmt.Sprintf(""))
		}
	}

	return output
}