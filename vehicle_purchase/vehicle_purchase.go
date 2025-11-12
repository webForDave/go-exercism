package purchase

func NeedLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	}
	return option2 + " is clearly the better choice."
} 

func CalculateResellPrice(originalPrice, age float64) float64 {
	percentagePrice := originalPrice / 100
	
	if age >= 10 {
		return percentagePrice * 50
	}else if age < 3 {
		return percentagePrice * 80
	} 
	return percentagePrice * 70
}