package lasagna

func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}

	return len(layers) * averagePreparationTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))

	for i := 0; i < len(quantities); i++ {
		scaledQuantities[i] = quantities[i] * float64(numberOfPortions) / 2
	}

	return scaledQuantities
}
