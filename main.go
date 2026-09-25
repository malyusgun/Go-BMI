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
	fmt.Printf("Your BMI is: %0.2f \n", BMI)

	switch {
	case BMI < 16:
		fmt.Println("You are significantly underweight")
	case BMI < 18.5:
		fmt.Println("You are underweight")
	case BMI < 25:
		fmt.Println("You have a normal weight")
	case BMI < 30:
		fmt.Println("You are overweight")
	case BMI < 35:
		fmt.Println("You have grade 1 obesity")
	case BMI < 40:
		fmt.Println("You have grade 2 obesity")
	default:
		fmt.Println("You have grade 3 obesity")
	}
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
