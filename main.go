package main

import (
	"fmt"
	"math"
)

const CENTIMETERS_PER_METER = 100
const HEIGHT_POWER = 2

func main() {
	fmt.Println("__BMI calculator__")

	userWeight, userHeight := getUserParams()
	BMI := userWeight / math.Pow(userHeight/CENTIMETERS_PER_METER, HEIGHT_POWER)

	outputResult(BMI)
}

func outputResult(BMI float64) {
	fmt.Printf("Your BMI is: %0.2f", BMI)
}

func getUserParams() (float64, float64) {
	var weight float64
	var height float64

	fmt.Print("Enter your weight: ")
	fmt.Scan(&weight)
	fmt.Print("Enter your height: ")
	fmt.Scan(&height)

	return weight, height
}
