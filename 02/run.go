package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Quest: BMI Calculator")

	// Recovering from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	fmt.Print("Enter your weight in kilos: ")
	weight, err := parse("weight")
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter your height in cm: ")
	height, err := parse("height")
	if err != nil {
		panic(err)
	}

	bmi := weight / math.Pow(height/100, 2)
	var bmiStatus string

	if bmi <= 15.0 {
		bmiStatus = "Very severely underweight"
	} else if bmi <= 16.0 {
		bmiStatus = "Severely underweight"
	} else if bmi <= 18.5 {
		bmiStatus = "Underweight"
	} else if bmi <= 25.0 {
		bmiStatus = "Normal"
	} else if bmi <= 30.0 {
		bmiStatus = "Overweight"
	} else if bmi <= 35.0 {
		bmiStatus = "Obese Class I"
	} else if bmi <= 40.0 {
		bmiStatus = "Obese Class II"
	} else {
		bmiStatus = "Obese Class III"
	}

	fmt.Printf("Your BMI is: %.2f (%s)", bmi, bmiStatus)
	fmt.Println()

}

func parse(name string) (float64, error) {
	var promt string
	_, err := fmt.Scanln(&promt)

	if err != nil {
		return 0, errors.New(fmt.Sprintf("Can't read your %s", name))
	}

	parsed, err := strconv.ParseFloat(promt, 64)

	if err != nil {
		return 0, errors.New(fmt.Sprintf("Can't parse your %s, probably it's not a float/numeric", name))
	}

	if parsed < 0 {
		return 0, errors.New(fmt.Sprintf("Your %s can't be less than 0", name))
	}

	return parsed, nil
}
