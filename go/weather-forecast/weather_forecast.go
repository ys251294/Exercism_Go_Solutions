// Package weather forecast weather depending upon city and condition.
package weather

// CurrentCondition represent current condition of city.
var CurrentCondition string

// CurrentLocation represent current location of city.
var CurrentLocation string

// Forecast return a string describing the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
