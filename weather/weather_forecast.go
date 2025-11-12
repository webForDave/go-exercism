//Package weather forecasts the weather condition of a given location.
package weather

// CurrentCondition holds the weather condition.
var CurrentCondition string

//CurrentLocation holds the location of the weather condition.
var CurrentLocation  string

/*
* Forecast takes a city and a weather condition and returns a readable string.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}