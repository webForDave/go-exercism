package collatzconjecture

import(
	"errors"
)

func CollatzConjecture(n int) (int, error){
	if n < 1{
		return 0, errors.New("invalid integer: number must be greater than or equal to one")
	}

	var numberOfSteps int

	for numberOfSteps = 0; n != 1; numberOfSteps ++{
		if n % 2 == 0{
			n = n/2
		}else if n % 2 != 0{
			n = (n * 3) + 1
		}

	}
	return numberOfSteps, nil
}