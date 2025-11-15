package lasagnamaster


func PreparationTime(layers []string, averagePreparationTimePerLayer int) int{
	if averagePreparationTimePerLayer == 0{
		return len(layers) * 2
	}
	return len(layers) * averagePreparationTimePerLayer
}

func Quantities(layers []string) (int, float64) {
	noodles, sauce := 0, 0.0
	for _, v := range layers{
		switch {
		case v == "noodles":
			noodles += 50
		case v == "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

func ScaleRecipe(amountForTwoPortions []float64, numberOfPortions int) []float64 {

	newAmount := make([]float64, len(amountForTwoPortions))

	copy(newAmount, amountForTwoPortions)

	for i, v := range newAmount{
		newAmount[i] = v * float64(numberOfPortions) / 2
	}

	return newAmount
}