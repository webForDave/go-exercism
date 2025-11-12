//Package cars analyzes the production in a car factory.
package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	newProductionRate := productionRate /  100
	println(newProductionRate)
	return int((float64(newProductionRate) * successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	costForOneCar, costForTenCars, overflow := 10000, 95000, 0

	// divides the carsCount argument by 10 and takes the remainder
	// if a remainder exists, it converts the remainder to an integer
	if carsCount % 10 != 0 {
		overflow = int(carsCount % 10)
	}

	// divides the number of cars by 10 and converts the the type of the result to integer
	numberOfCars := int(carsCount / 10)

	cost := (numberOfCars * costForTenCars) + (overflow * costForOneCar)

	return uint(cost)
}