package main

func main() {
}
func fuelDistance(fuelConsumption int) int {
	const Kilometers = 100
	fuelConsumption = 10
	fuelAvailable := 20
	Distance := (fuelAvailable * Kilometers) / fuelConsumption
	return Distance
}
