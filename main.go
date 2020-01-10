package main

func main() {
}
func fuelDistance(fuelConsumption, fuelAvailable int) int {
	const kilometers = 100
	distance := (fuelAvailable * kilometers) / fuelConsumption
	return distance
}
