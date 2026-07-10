// Package weather provides tools to forecast the current weather condition.
package weather

var (
// CurrentCondition represents the current weather condition.
	CurrentCondition string

// CurrentLocation represents the current location.
	CurrentLocation  string
)

// Forecast returns a string with the weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
