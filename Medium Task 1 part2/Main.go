package main

import (
	"fmt"
	"math"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {
	var amSpeed AmericanVelocity = 130
	amSpeed = amSpeed / 1609 * 3600
	amSpeed = AmericanVelocity(math.Round(float64(amSpeed)*100) / 100)
	var euSpeed EuropeanVelocity = 120.4
	euSpeed = euSpeed / 1000 * 3600
	euSpeed = EuropeanVelocity(math.Round(float64(euSpeed)*100) / 100)
	fmt.Println(amSpeed)
	fmt.Println(euSpeed)

}
