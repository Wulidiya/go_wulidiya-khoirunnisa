package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c Car) CalculateDistance() float64 {
	fuelEfficiency := 1.5                       // satuan liter/mill
	distancePerLiter := 1000.0 / fuelEfficiency // satuan meter
	distance := distancePerLiter * c.FuelIn
	return distance
}

func main() {
	myCar := Car{
		Type:   "Sedan",
		FuelIn: 20.0,
	}

	distance := myCar.CalculateDistance()
	fmt.Printf("%s dapat menempuh jarak kurang lebih %.2f meter dengan bahan bakar saat ini.\n", myCar.Type, distance)
}
