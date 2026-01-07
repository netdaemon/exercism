// Package weather provides tools to return the current condition for the location.
package weather

var (
    // CurrentCondition stores the current condition for the location.
	CurrentCondition string
    // CurrentLocation stores the current location.
	CurrentLocation  string
)

// Forecast returns a string value showing information from the passed in city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition

	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
